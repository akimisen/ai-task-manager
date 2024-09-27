package main

import (
    "ai-task-manager/internal/api"
    "ai-task-manager/internal/queue"
    "ai-task-manager/internal/services"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "log"
    "net/http"
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
    ttsService := services.NewTTSService(memoryQueue)

    router := gin.Default()
    api.SetupRoutes(router, ttsService)

    log.Fatal(http.ListenAndServe(":8080", router))
}