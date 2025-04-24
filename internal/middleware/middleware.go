/**
 * Package repo
 * @file      : middleware.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/23 12:12
 **/

package middleware

import (
	"github.com/volcengine/skd/internal/router"
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) {
		start := time.Now()
		log.Printf("[%s] %s | Started\n", ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
		log.Printf("[%s] %s | Completed in %v\n", ctx.Request.Method, ctx.Request.URL.Path, time.Since(start))
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) {
		token := ctx.Request.Header.Get("Authorization")
		log.Printf("[%s] %s | %s | Auth\n", ctx.Request.Method, ctx.Request.URL.Path, token)
		// TODO 认证逻辑
		//if token != "Bearer valid-token" {
		//	ctx.Error(volc.StatusUnauthorized, "Unauthorized")
		//	return // 中断执行
		//}
		next(ctx)
	}
}

func CorsMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) {
		log.Printf("[%s] %s | Cors\n", ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		next(ctx)
	}
}

func RecoveryMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) {
		log.Printf("[%s] %s | Recovery\n", ctx.Request.Method, ctx.Request.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				ctx.Error(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		next(ctx)
	}
}
