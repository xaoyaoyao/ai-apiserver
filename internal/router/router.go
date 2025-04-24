/**
 * Package repo
 * @file      : router.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 16:38
 **/

package router

import (
	"encoding/json"
	"net/http"
)

// Context 请求上下文，封装请求和响应
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// HandlerFunc 处理函数类型
type HandlerFunc func(*Context)

// Middleware 中间件类型（函数式）
type Middleware func(HandlerFunc) HandlerFunc

// Router 路由结构（支持中间件链和路由分组）
type Router struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

// NewRouter 创建新路由
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// Use 添加全局中间件
func (r *Router) Use(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// Group 创建路由分组（继承父级中间件）
func (r *Router) Group(prefix string, middlewares ...Middleware) *Group {
	return &Group{
		router:      r,
		prefix:      prefix,
		middlewares: append(r.middlewares, middlewares...), // 继承全局中间件
	}
}

// Handle 注册路由（核心方法）
func (r *Router) Handle(method, path string, handler HandlerFunc, middlewares ...Middleware) {
	// 合并全局中间件和路由级中间件
	allMiddlewares := append(r.middlewares, middlewares...)

	// 构建中间件链
	finalHandler := handler
	for i := len(allMiddlewares) - 1; i >= 0; i-- {
		finalHandler = allMiddlewares[i](finalHandler)
	}

	// 包装为 volc.HandlerFunc
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ctx := &Context{Writer: w, Request: req}
		finalHandler(ctx)
	})
}

// ServeHTTP 实现 http.Handler 接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

type Group struct {
	router      *Router
	prefix      string
	middlewares []Middleware
}

func (g *Group) Handle(method, path string, handler HandlerFunc, middlewares ...Middleware) {
	fullPath := g.prefix + path
	allMiddlewares := append(g.middlewares, middlewares...)
	g.router.Handle(method, fullPath, handler, allMiddlewares...)
}

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

func HealthCheck(ctx *Context) {
	ctx.makeOK(nil)
}
