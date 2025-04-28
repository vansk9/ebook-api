package service

import (
	"context"
	"ebook-api/dto"
	"ebook-api/models"
	"ebook-api/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (dto.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]dto.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (dto.UserResponse, error) {
	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:    createdUser.UserID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID:    user.UserID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return responses, nil
}
