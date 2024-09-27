package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis(addr string) {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// 测试连接
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
