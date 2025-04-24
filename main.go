package main

import (
	"github.com/volcengine/skd/internal/config"
	"github.com/volcengine/skd/internal/middleware"
	"github.com/volcengine/skd/internal/router"
	"log"
	"net/http"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 初始化
	_ = config.Init()

	mux := router.NewRouter()

	// 全局中间件（所有路由生效）
	mux.Use(middleware.LoggerMiddleware)

	// 路由分组示例
	apiGroup := mux.Group("/api")
	{
		// 分组级中间件（仅该分组生效）
		apiGroup.Handle("GET", "/health", router.HealthCheck)
	}

	mux.Use(middleware.CorsMiddleware)
	mux.Use(middleware.RecoveryMiddleware)

	// 带认证的路由
	authGroup := mux.Group("/api/ai", middleware.AuthMiddleware)
	{
		authGroup.Handle("POST", "/v1/volcengine", router.ProcessVolcEngine)
		authGroup.Handle("POST", "/v1/meitu", router.ProcessMeitu)
	}

	// 启动服务器配置
	server := &http.Server{
		Addr:         config.Get().Addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
