package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akimisen/ai-task-manager/internal/models"
	"github.com/redis/go-redis/v9"
)

type ImagenService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewImagenService(redisClient *redis.Client) *ImagenService {
	return &ImagenService{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *ImagenService) CreateTask(task *models.ImagenTask) error {
	taskID := fmt.Sprintf("imagen:task:%d", task.ID)
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = s.redisClient.Set(s.ctx, taskID, taskData, 0).Err()
	if err != nil {
		return err
	}

	return s.redisClient.LPush(s.ctx, "imagen_task_queue", taskID).Err()
}

func (s *ImagenService) GetTask(taskID string) (*models.ImagenTask, error) {
	taskData, err := s.redisClient.Get(s.ctx, taskID).Result()
	if err != nil {
		return nil, err
	}

	var task models.ImagenTask
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *ImagenService) ListTasks() ([]*models.ImagenTask, error) {
	taskIDs, err := s.redisClient.LRange(s.ctx, "imagen_task_queue", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var tasks []*models.ImagenTask
	for _, taskID := range taskIDs {
		task, err := s.GetTask(taskID)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
