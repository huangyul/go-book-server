package repository

import (
	"context"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
