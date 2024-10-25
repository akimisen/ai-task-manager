package repository

import "ai-task-manager/internal/model"

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id string) (*model.Task, error)
	Update(task *model.Task) error
	List() ([]model.Task, error)
}

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Update(user *model.User) error
}
