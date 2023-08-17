package repository

import (
	"context"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
var ErrNotFound = dao.ErrNotFound

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	user, err := r.dao.FindByEmail(ctx, u)
	ud := domain.User{
		ID:       user.Id,
		Password: user.Password,
		Email:    user.Email,
	}
	return ud, err
}
