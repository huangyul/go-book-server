package service

import (
	"context"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// SignUp 注册
func (usc *UserService) SignUp(ctx context.Context, user domain.User) error {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return usc.repo.Create(ctx, user)
}
