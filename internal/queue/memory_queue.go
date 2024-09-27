package queue

import (
    "ai-task-manager/internal/models"
    "sync"
)

type MemoryQueue struct {
    tasks []models.Task
    mutex sync.Mutex
}

func NewMemoryQueue() *MemoryQueue {
    return &MemoryQueue{
        tasks: make([]models.Task, 0),
    }
}

// NewMemoryQueueWithMockData 创建一个包含模拟数据的 MemoryQueue
func NewMemoryQueueWithMockData() *MemoryQueue {
    queue := NewMemoryQueue()
    mockTasks := []models.Task{
        &models.TTSTask{
            ID:     "1",
            Type:   models.TaskTypeTTS,
            Status: "pending",
            Text:   "Hello, this is a test TTS task.",
        },
        &models.TTSTask{
            ID:     "2",
            Type:   models.TaskTypeTTS,
            Status: "processing",
            Text:   "Another test task for TTS processing.",
        },
        &models.TTSTask{
            ID:     "3",
            Type:   models.TaskTypeTTS,
            Status: "completed",
            Text:   "This task has been completed.",
        },
    }

    for _, task := range mockTasks {
        queue.Push(task)
    }

    return queue
}

func (q *MemoryQueue) Push(task models.Task) error {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    q.tasks = append(q.tasks, task)
    return nil
}

func (q *MemoryQueue) Pop() (models.Task, error) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    if len(q.tasks) == 0 {
        return nil, nil
    }
    task := q.tasks[0]
    q.tasks = q.tasks[1:]
    return task, nil
}

func (q *MemoryQueue) List() ([]models.Task, error) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    return append([]models.Task{}, q.tasks...), nil
}