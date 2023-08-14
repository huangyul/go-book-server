package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	// 获取毫秒数
	now := time.Now().UnixMilli()
	u.CreateTime = now
	u.UpdateTime = now
	return dao.db.WithContext(ctx).Create(&u).Error
}

// User 数据库表结构
type User struct {
	Id       int64
	Email    string
	Password string

	CreateTime int64
	UpdateTime int64
}
