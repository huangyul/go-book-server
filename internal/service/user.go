package service

import (
	"context"
	"errors"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("用户/邮箱或密码不对")

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
	err = usc.repo.Create(ctx, user)

	return err
}

func (usc *UserService) FindByEmail(ctx context.Context, user domain.User) (domain.User, error) {
	u, err := usc.repo.FindByEmail(ctx, user)

	if errors.Is(err, repository.ErrNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}

	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (usc *UserService) UpdateUserInfo(ctx context.Context, user domain.User) error {
	return usc.repo.UpdateUser(ctx, user)
}

func (usc *UserService) GetUserInfoById(ctx context.Context, user domain.User) (domain.User, error) {
	return usc.repo.GetUserInfo(ctx, user)
}

func (usc *UserService) profile(ctx context.Context, id int64) (domain.User, err) {
	return usc.profile(ctx, id)
}
