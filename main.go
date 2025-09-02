package main

import (
	"fmt"
	"log"
	"os"

	"gin-auth-project/config"
	"gin-auth-project/database"
	"gin-auth-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.Init()

	// 设置Gin模式
	gin.SetMode(config.AppConfig.ServerMode)

	// 初始化数据库连接
	database.InitPostgres()

	// 初始化Redis连接
	database.InitRedis()

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	port := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
	log.Printf("Server starting on port %s", port)
	log.Printf("Environment: %s", config.AppConfig.ServerMode)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
		os.Exit(1)
	}
}
