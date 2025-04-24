package main

import (
	"github.com/volcengine/skd/internal/common/endpoint"
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
	apiGroup := mux.Group(endpoint.ROOT_PATH)
	{
		// 分组级中间件（仅该分组生效）
		apiGroup.Handle(endpoint.METHOD_GET, endpoint.HEALTH_PATH, router.HealthCheck)
	}

	mux.Use(middleware.CorsMiddleware)
	mux.Use(middleware.RecoveryMiddleware)

	// 带认证的路由
	authGroup := mux.Group(endpoint.API_STSTEM_PATH, middleware.AuthMiddleware)
	{
		authGroup.Handle(endpoint.METHOD_GET, endpoint.USERS_POSTS_PATH, router.ProcessUserPosts)
	}

	// 带认证的AI路由
	aiGroup := mux.Group(endpoint.API_AI_PATH, middleware.AuthMiddleware)
	{
		aiGroup.Handle(endpoint.METHOD_POST, endpoint.VOLC_ENGINE_PATH, router.ProcessVolcEngine)
		aiGroup.Handle(endpoint.METHOD_POST, endpoint.MEITU_PATH, router.ProcessMeitu)
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
