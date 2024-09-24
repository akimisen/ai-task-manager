package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akimisen/ai-task-manager/internal/models"
	"github.com/redis/go-redis/v9"
)

type TTSService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewTTSService(redisClient *redis.Client) *TTSService {
	return &TTSService{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *TTSService) CreateTask(task *models.TTSTask) error {
	taskID := fmt.Sprintf("tts:task:%d", task.ID)
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = s.redisClient.Set(s.ctx, taskID, taskData, 0).Err()
	if err != nil {
		return err
	}

	return s.redisClient.LPush(s.ctx, "tts_task_queue", taskID).Err()
}

func (s *TTSService) GetTask(taskID string) (*models.TTSTask, error) {
	taskData, err := s.redisClient.Get(s.ctx, taskID).Result()
	if err != nil {
		return nil, err
	}

	var task models.TTSTask
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TTSService) ListTasks() ([]*models.TTSTask, error) {
	taskIDs, err := s.redisClient.LRange(s.ctx, "tts_task_queue", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var tasks []*models.TTSTask
	for _, taskID := range taskIDs {
		task, err := s.GetTask(taskID)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
