package api

import (
	"ai-task-manager/internal/api/handlers"
	"ai-task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, ttsService *services.TTSService) {
	ttsHandler := handlers.NewTTSHandler(ttsService)

	v1 := router.Group("/api/v1")
	{
		tts := v1.Group("/tts")
		{
			tts.POST("/tasks", ttsHandler.CreateTask)
			tts.GET("/tasks/:id", ttsHandler.GetTask)
			tts.GET("/tasks/:id/status", ttsHandler.GetTaskStatus)
			tts.GET("/tasks", ttsHandler.ListTasks)
		}
	}
}
