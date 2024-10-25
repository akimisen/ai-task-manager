package main

import (
	"ai-task-manager/internal/api"
	"ai-task-manager/internal/queue"
	"ai-task-manager/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	gin.SetMode(gin.DebugMode)
	memoryQueue := queue.NewMemoryQueueWithMockData()
	ttsService := service.NewTTSService(memoryQueue)

	router := gin.Default()
	api.SetupRoutes(router, ttsService)

	log.Fatal(http.ListenAndServe(":8080", router))
}
