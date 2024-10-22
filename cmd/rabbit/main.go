package main

import (
	"ai-task-manager/internal/config"
	"ai-task-manager/internal/repository"
	"ai-task-manager/pkg/rabbitmq"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	rmq, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	repo := repository.NewRabbitMQRepository(rmq)
	svc := service.NewTTSService(repo)
	h := handler.NewTTSHandler(svc)

	r := gin.Default()
	r.POST("/tts", h.CreateTTSTask)

	go svc.ConsumeTask()

	r.Run(cfg.Server.Address)
}
