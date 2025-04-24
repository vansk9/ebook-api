package service

import (
	"context"
	"ebook-api/models"
	"ebook-api/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}
