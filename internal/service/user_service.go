package service

import (
	"ai-task-manager/internal/model"
	"ai-task-manager/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

// 可以根据需要添加更多方法
