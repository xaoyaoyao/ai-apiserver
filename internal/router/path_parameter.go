/**
 * Package repo
 * @file      : path_parameter.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 12:47
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

// 绑定路径参数到结构体
func BindPathParams(r *http.Request, pattern string, dest interface{}) error {
	// 1. 路径清理
	cleanPath := strings.Split(r.URL.Path, "?")[0] // 移除查询参数
	cleanPath = strings.Split(cleanPath, "#")[0]   // 移除锚点
	cleanPath = strings.TrimRight(cleanPath, "/")  // 移除末尾斜杠

	// 2. 编译正则表达式
	re, paramNames, err := compilePattern(pattern)
	if err != nil {
		return fmt.Errorf("pattern compile failed: %w", err)
	}

	// 3. 执行路径匹配
	matches := re.FindStringSubmatch(cleanPath)
	if matches == nil {
		return &PathMismatchError{
			URL:       r.URL.Path,
			Pattern:   pattern,
			CleanPath: cleanPath,
			ParamMap:  buildParamMap(paramNames, matches),
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
		paramValue, exists := matchesMap(paramNames, matches, tag)
		if !exists {
			return &MissingParamError{Param: tag}
		}

		// 类型安全转换
		if err := convertAndSet(field, paramValue); err != nil {
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
func compilePattern(pattern string) (*regexp.Regexp, []string, error) {
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
func matchesMap(paramNames []string, matches []string, targetTag string) (string, bool) {
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
func convertAndSet(field reflect.Value, value string) error {
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

func buildParamMap(paramNames []string, matches []string) map[string]string {
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

// 辅助函数
func paramValues(matches []string) []string {
	values := make([]string, len(matches))
	copy(values, matches[1:]) // 跳过全匹配组
	return values
}

// 错误处理
func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	switch e := err.(type) {
	case *PathMismatchError:
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Path mismatch",
			"details": e.ParamMap,
		})
	case *MissingParamError:
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Missing parameter",
			"param": e.Param,
		})
	case *ConversionError:
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Type conversion error",
			"field":    e.Field,
			"input":    e.Input,
			"expected": e.Detail,
		})
	default:
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
	}
}
