package queue

import "ai-task-manager/internal/models"

type TaskQueue interface {
    Push(task models.Task) error
    Pop() (models.Task, error)
    List() ([]models.Task, error)
}