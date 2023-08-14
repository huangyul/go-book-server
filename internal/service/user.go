package service

import (
	"context"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (usc *UserService) SignUp(ctx context.Context, user domain.User) error {
	return nil
}
