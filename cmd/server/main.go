package main

import (
    "ai-task-manager/internal/api"
    "ai-task-manager/internal/config"
    "ai-task-manager/internal/queue"
    "ai-task-manager/internal/services"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "log"
    "net/http"
)

func main() {
    cfg := config.Load()

    redisClient := redis.NewClient(&redis.Options{
        Addr:     cfg.RedisAddr,
        Password: cfg.RedisPassword,
        DB:       0,
    })

    redisQueue := queue.NewRedisQueue(redisClient, "tts_task_queue")
    ttsService := services.NewTTSService(redisQueue)

    router := gin.Default()
    api.SetupRoutes(router, ttsService)

    log.Fatal(http.ListenAndServe(":8080", router))
}