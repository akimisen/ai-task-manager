package api

import (
	"your_project/internal/api/handlers"
	"your_project/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(router *gin.Engine) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
		DB:   0,                // 使用默认数据库
	})

	// 创建服务实例
	imagenService := services.NewImagenService(redisClient)
	ttsService := services.NewTTSService(redisClient)
	videoEditingService := services.NewVideoEditingService(redisClient)

	// 创建处理器实例
	imagenHandler := handlers.NewImagenHandler(imagenService)
	ttsHandler := handlers.NewTTSHandler(ttsService)
	videoEditingHandler := handlers.NewVideoEditingHandler(videoEditingService)

	v1 := router.Group("/api/v1")
	{
		// Imagen task routes
		imagen := v1.Group("/imagen")
		{
			imagen.POST("/tasks", imagenHandler.CreateTask)
			imagen.GET("/tasks/:id", imagenHandler.GetTask)
			imagen.GET("/tasks", imagenHandler.ListTasks)
		}

		// TTS task routes
		tts := v1.Group("/tts")
		{
			tts.POST("/tasks", ttsHandler.CreateTask)
			tts.GET("/tasks/:id", ttsHandler.GetTask)
			tts.GET("/tasks", ttsHandler.ListTasks)
		}

		// Video Editing task routes
		videoEditing := v1.Group("/video_editing")
		{
			videoEditing.POST("/tasks", videoEditingHandler.CreateTask)
			videoEditing.GET("/tasks/:id", videoEditingHandler.GetTask)
			videoEditing.GET("/tasks", videoEditingHandler.ListTasks)
		}
	}
}
