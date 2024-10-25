package queue

import (
	"ai-task-manager/internal/model"
	"fmt"
)

type TaskQueue interface {
	Publish(task model.Task) error
	Consume() (<-chan model.Task, error)
	Close() error
}

// NewTaskQueue 创建一个新的任务队列
func NewTaskQueue(queueType string, config map[string]string) (TaskQueue, error) {
	switch queueType {
	case "rabbitmq":
		return NewRabbitMQ(config["url"])
	// 未来可以添加其他类型的队列
	// case "memory":
	//     return NewMemoryQueue()
	// case "redis":
	//     return NewRedisQueue(config["address"])
	default:
		return nil, fmt.Errorf("unsupported queue type: %s", queueType)
	}
}
