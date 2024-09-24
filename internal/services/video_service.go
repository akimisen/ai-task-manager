package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akimisen/ai-task-manager/internal/models"

	"github.com/redis/go-redis/v9"
)

type VideoService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewVideoService(redisClient *redis.Client) *VideoService {
	return &VideoService{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *VideoService) CreateTask(task *models.Task) error {
	taskID := fmt.Sprintf("video_editing:task:%d", task.ID)
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = s.redisClient.Set(s.ctx, taskID, taskData, 0).Err()
	if err != nil {
		return err
	}

	return s.redisClient.LPush(s.ctx, "video_task_queue", taskID).Err()
}

func (s *VideoService) GetTask(taskID string) (*models.Task, error) {
	taskData, err := s.redisClient.Get(s.ctx, taskID).Result()
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *VideoService) ListTasks() ([]*models.Task, error) {
	taskIDs, err := s.redisClient.LRange(s.ctx, "video_task_queue", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, taskID := range taskIDs {
		task, err := s.GetTask(taskID)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
