package main

import (
	"ai-task-manager/internal/config"
	"ai-task-manager/internal/handler"
	"ai-task-manager/internal/repository"
	"ai-task-manager/internal/service"
	"ai-task-manager/pkg/logger"
	"ai-task-manager/pkg/queue"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志系统
	logger := logger.NewLogger()
	defer logger.Sync() // 确保所有日志都被写入

	// 加载配置文件
	cfg := config.Load()
	logger.Info("Configuration loaded successfully")

	// 初始化 RabbitMQ 连接
	queueConfig := map[string]string{
		"url": cfg.RabbitMQ.URL,
	}
	taskQueue, err := queue.NewTaskQueue("rabbitmq", queueConfig)
	if err != nil {
		logger.Fatalf("Failed to create task queue: %v", err)
	}
	defer taskQueue.Close()
	logger.Info("RabbitMQ connection established")

	// 初始化存储库
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewRabbitMQRepository(taskQueue)

	// 初始化服务
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	ttsService := service.NewTTSService(taskRepo)

	// 设置路由
	router := gin.Default()
	handler.SetupRoutes(router, authService, ttsService, userService)

	// 启动后台任务消费者
	go ttsService.ConsumeTask()
	logger.Info("Background task consumer started")

	// 启动 HTTP 服务器
	logger.Infof("Starting server on %s", cfg.Server.Address)
	if err := router.Run(cfg.Server.Address); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
