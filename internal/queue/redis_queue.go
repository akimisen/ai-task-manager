package queue

import (
    "ai-task-manager/internal/models"
    "context"
    "encoding/json"
    "fmt"
    "github.com/redis/go-redis/v9"
)

type RedisQueue struct {
    client *redis.Client
    key    string
}

func NewRedisQueue(client *redis.Client, key string) *RedisQueue {
    return &RedisQueue{
        client: client,
        key:    key,
    }
}

func (q *RedisQueue) Push(task models.Task) error {
    taskData, err := json.Marshal(struct {
        Type string      `json:"type"`
        Data models.Task `json:"data"`
    }{
        Type: string(task.GetType()),
        Data: task,
    })
    if err != nil {
        return err
    }
    return q.client.RPush(context.Background(), q.key, taskData).Err()
}

func (q *RedisQueue) Pop() (models.Task, error) {
    taskData, err := q.client.LPop(context.Background(), q.key).Bytes()
    if err == redis.Nil {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return q.unmarshalTask(taskData)
}

func (q *RedisQueue) List() ([]models.Task, error) {
    taskDataList, err := q.client.LRange(context.Background(), q.key, 0, -1).Result()
    if err != nil {
        return nil, err
    }
    var tasks []models.Task
    for _, taskData := range taskDataList {
        task, err := q.unmarshalTask([]byte(taskData))
        if err != nil {
            continue
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (q *RedisQueue) unmarshalTask(data []byte) (models.Task, error) {
    var taskWrapper struct {
        Type string          `json:"type"`
        Data json.RawMessage `json:"data"`
    }
    err := json.Unmarshal(data, &taskWrapper)
    if err != nil {
        return nil, err
    }

    var task models.Task
    switch models.TaskType(taskWrapper.Type) {
    case models.TaskTypeTTS:
        task = &models.TTSTask{}
    // case models.TaskTypeImageGen:
    //     task = &models.ImageGenTask{}
    default:
        return nil, fmt.Errorf("unknown task type: %s", taskWrapper.Type)
    }

    err = json.Unmarshal(taskWrapper.Data, task)
    if err != nil {
        return nil, err
    }

    return task, nil
}