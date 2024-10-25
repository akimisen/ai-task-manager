package handler

import (
	"ai-task-manager/internal/handler/auth"
	"ai-task-manager/internal/handler/task"
	"ai-task-manager/internal/handler/user"
	"ai-task-manager/internal/middleware"
	"ai-task-manager/internal/service"
	"ai-task-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, log *logger.Logger, authService *service.AuthService, ttsService *service.TTSService, userService *service.UserService) {
	// 使用日志中间件
	router.Use(middleware.LoggerMiddleware(log))

	// 初始化各种 handler
	authHandler := auth.NewAuthHandler(authService)
	ttsHandler := task.NewTTSHandler(ttsService)
	userHandler := user.NewUserHandler(userService)

	// API 版本组
	v1 := router.Group("/api/v1")

	// 公开路由
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	// 需要认证的路由
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户相关路由
		user := protected.Group("/users")
		{
			user.GET("/me", userHandler.GetUserInfo)
			user.PUT("/me", userHandler.UpdateUserInfo)
		}

		// TTS 任务相关路由
		tts := protected.Group("/tts")
		{
			tts.POST("/tasks", ttsHandler.CreateTask)
			tts.GET("/tasks/:id", ttsHandler.GetTask)
			tts.GET("/tasks/:id/status", ttsHandler.GetTaskStatus)
			tts.GET("/tasks", ttsHandler.ListTasks)
		}

		// 可以在这里添加其他任务类型的路由，如图像生成等
	}
}
