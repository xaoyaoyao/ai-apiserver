/**
 * Package repo
 * @file      : router_parameter.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 14:29
 **/

package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Query 获取 URL 查询参数
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// PostForm 获取表单参数
func (c *Context) PostForm(key string) (string, error) {
	err := c.Request.ParseForm()
	if err != nil {
		return "", err
	}
	return c.Request.FormValue(key), nil
}

// BindJSON 解析 JSON 请求体
func (c *Context) BindJSON(v interface{}) error {
	defer c.Request.Body.Close()
	return json.NewDecoder(c.Request.Body).Decode(v)
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
}

// JSON 返回 JSON 响应
func (c *Context) JSON(statusCode int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)
	json.NewEncoder(c.Writer).Encode(data)
}

// Error 返回错误响应
func (c *Context) Error(statusCode int, message string) {
	c.JSON(statusCode, map[string]string{"error": message})
}

func (c *Context) makeOK(data interface{}) {
	c.JSON(http.StatusOK, c.makeData(http.StatusOK, "OK", data))
}

func (c *Context) makeData(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (c *Context) makeError(statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// 绑定路径参数到结构体
func (c *Context) BindPathParams(pattern string, dest interface{}) error {
	// 1. 路径清理
	cleanPath := strings.Split(c.Request.URL.Path, "?")[0] // 移除查询参数
	cleanPath = strings.Split(cleanPath, "#")[0]           // 移除锚点
	cleanPath = strings.TrimRight(cleanPath, "/")          // 移除末尾斜杠

	// 2. 编译正则表达式
	re, paramNames, err := c.compilePattern(pattern)
	if err != nil {
		return fmt.Errorf("pattern compile failed: %w", err)
	}

	// 3. 执行路径匹配
	matches := re.FindStringSubmatch(cleanPath)
	if matches == nil {
		return &PathMismatchError{
			URL:       c.Request.URL.Path,
			Pattern:   pattern,
			CleanPath: cleanPath,
			ParamMap:  c.buildParamMap(paramNames, matches),
		}
	}

	// 4. 反射绑定到结构体
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.IsNil() {
		return fmt.Errorf("dest must be non-nil pointer to struct")
	}
	destValue = destValue.Elem()
	destType := destValue.Type()

	// 5. 遍历结构体字段进行绑定
	for i := 0; i < destValue.NumField(); i++ {
		field := destValue.Field(i)
		fieldType := destType.Field(i)
		tag := fieldType.Tag.Get("path")

		if tag == "" {
			continue
		}

		// 获取参数值
		paramValue, exists := c.matchesMap(paramNames, matches, tag)
		if !exists {
			return &MissingParamError{Param: tag}
		}

		// 类型安全转换
		if err := c.convertAndSet(field, paramValue); err != nil {
			return &ConversionError{
				Field:  fieldType.Name,
				Tag:    tag,
				Input:  paramValue,
				Detail: err.Error(),
			}
		}
	}

	return nil
}

// 路径模式编译（增强版）
func (c *Context) compilePattern(pattern string) (*regexp.Regexp, []string, error) {
	var paramNames []string
	parts := strings.Split(pattern, "/")

	reParts := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := part[1 : len(part)-1]
			paramNames = append(paramNames, paramName)
			reParts = append(reParts, "([^/]+)")
		} else {
			reParts = append(reParts, regexp.QuoteMeta(part))
		}
	}

	reStr := "^/" + strings.Join(reParts, "/") + "/?$" // 允许末尾斜杠
	re, err := regexp.Compile(reStr)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	return re, paramNames, nil
}

// 辅助函数：构建参数映射
func (c *Context) matchesMap(paramNames []string, matches []string, targetTag string) (string, bool) {
	for i, name := range paramNames {
		if name == targetTag {
			if i+1 < len(matches) {
				return matches[i+1], true
			}
			return "", false
		}
	}
	return "", false
}

// 类型安全转换
func (c *Context) convertAndSet(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %q", value)
		}
		field.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid unsigned integer: %q", value)
		}
		field.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid float: %q", value)
		}
		field.SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean: %q", value)
		}
		field.SetBool(b)
	default:
		return fmt.Errorf("unsupported type: %s", field.Kind())
	}
	return nil
}

func (c *Context) buildParamMap(paramNames []string, matches []string) map[string]string {
	paramMap := make(map[string]string)
	for i, name := range paramNames {
		if i+1 < len(matches) {
			paramMap[name] = matches[i+1]
		} else {
			paramMap[name] = "<MISSING>"
		}
	}
	return paramMap
}

type PathMismatchError struct {
	URL       string
	Pattern   string
	ParamMap  map[string]string
	CleanPath string
}

func (e *PathMismatchError) Error() string {
	return fmt.Sprintf("path mismatch: %q does not match pattern %q (cleaned: %q)",
		e.URL, e.Pattern, e.CleanPath)
}

type MissingParamError struct {
	Param string
}

func (e *MissingParamError) Error() string {
	return fmt.Sprintf("missing required path parameter: %s", e.Param)
}

type ConversionError struct {
	Field  string
	Tag    string
	Input  string
	Detail string
}

func (e *ConversionError) Error() string {
	return fmt.Sprintf("conversion error for field %s (tag %s): %s (input: %q)",
		e.Field, e.Tag, e.Detail, e.Input)
}

// 错误处理
func (c *Context) HandleError(err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusBadRequest)

	switch e := err.(type) {
	case *PathMismatchError:
		json.NewEncoder(c.Writer).Encode(map[string]interface{}{
			"error":   "Path mismatch",
			"details": e.ParamMap,
		})
	case *MissingParamError:
		json.NewEncoder(c.Writer).Encode(map[string]string{
			"error": "Missing parameter",
			"param": e.Param,
		})
	case *ConversionError:
		json.NewEncoder(c.Writer).Encode(map[string]string{
			"error":    "Type conversion error",
			"field":    e.Field,
			"input":    e.Input,
			"expected": e.Detail,
		})
	default:
		json.NewEncoder(c.Writer).Encode(map[string]string{
			"error": "Internal server error",
		})
	}
}
