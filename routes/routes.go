package routes

import (
	"gin-auth-project/handlers"
	"gin-auth-project/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 添加CORS中间件
	r.Use(middleware.CORSMiddleware())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Server is running"})
	})

	// 认证相关路由（无需认证）
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.AuthHandler{}.Login)
		auth.POST("/register", handlers.AuthHandler{}.Register)
	}

	// 需要认证的路由
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 用户认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/logout", handlers.AuthHandler{}.Logout)
			auth.GET("/profile", handlers.AuthHandler{}.GetProfile)
			auth.PUT("/profile", handlers.AuthHandler{}.UpdateProfile)
			auth.POST("/refresh", handlers.AuthHandler{}.RefreshToken)
		}

		// 用户管理（需要管理员权限）
		users := api.Group("/users")
		users.Use(middleware.AdminMiddleware())
		{
			users.GET("", handlers.UserHandler{}.GetAllUsers)
			users.POST("", handlers.UserHandler{}.CreateUser)
			users.GET("/:id", handlers.UserHandler{}.GetUserByID)
			users.PUT("/:id", handlers.UserHandler{}.UpdateUser)
			users.DELETE("/:id", handlers.UserHandler{}.DeleteUser)
			users.PATCH("/:id/status", handlers.UserHandler{}.ToggleUserStatus)
		}

		// 受保护的资源路由（需要用户权限）
		protected := api.Group("/protected")
		protected.Use(middleware.UserMiddleware())
		{
			protected.GET("/data", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "This is protected data",
					"user_id": middleware.GetCurrentUserID(c),
					"role":    middleware.GetCurrentUserRole(c),
				})
			})
		}
	}

	return r
}
