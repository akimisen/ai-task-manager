package main

import (
	"log"
	"net/http"

	"github.com/akimisen/ai-task-manager/internal/api"
	// "github.com/akimisen/ai-task-manager/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg := config.Load()

	// 初始化 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,     // 假设配置中有 Redis 地址
		Password: cfg.RedisPassword, // 如果有密码的话
		DB:       0,                 // 使用默认 DB
	})

	// 可以在这里添加一个 Ping 操作来确保 Redis 连接正常
	// ctx := context.Background()
	// _, err := rdb.Ping(ctx).Result()
	// if err != nil {
	//     log.Fatal("无法连接到 Redis:", err)
	// }

	router := gin.Default()

	// 修改 SetupRoutes 函数，传入 redis.Client 而不是 database
	api.SetupRoutes(router, rdb)

	log.Fatal(http.ListenAndServe(":8080", router))
}
