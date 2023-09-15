package repository

import (
	"context"
	"go-book-server/internal/domain"
	"go-book-server/internal/repository/cache"
	"go-book-server/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
var ErrNotFound = dao.ErrNotFound

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: c,
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

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.cache.Get(ctx, id)
	if err == nil {
		// 如果err==nil，说明一定有数据
		return u, nil
	}
	// 现在要不就没数据，要不就是redis出错的
	// 这时候要考虑如果是redis崩了，继续向mysql查数据的话，大量的请求会把mysql也弄蹦，这时候就要做数据库限流
	// 也可以当redis出错时，就返回错误，不再找数据库

	// 下面使用第一种方式
	ue, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	u = domain.User{
		ID:       ue.Id,
		Email:    ue.Email,
		Password: ue.Password,
		NickName: ue.NickName,
		Birthday: ue.Birthday,
		Brief:    ue.Brief,
	}

	go func() {
		err = r.cache.Set(ctx, u)
		if err != nil {
			// 打日志
		}
	}()

	return u, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, u domain.User) error {
	err := r.dao.UpdateUserInfo(ctx, u)
	return err
}

func (r *UserRepository) GetUserInfo(ctx context.Context, u domain.User) (domain.User, error) {
	return r.dao.GetById(ctx, u)
}
