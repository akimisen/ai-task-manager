package service

import (
	"ai-task-manager/internal/model"
	"ai-task-manager/internal/repository"
	"ai-task-manager/pkg/auth"
	"errors"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(username, password, email string) error {
	// 检查用户是否已存在
	if _, err := s.userRepo.GetByUsername(username); err == nil {
		return errors.New("username already exists")
	}

	// 创建新用户
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &model.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	return s.userRepo.Create(newUser)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// 生成 JWT token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
