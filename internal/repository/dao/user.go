package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go-book-server/internal/domain"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrNotFound           = gorm.ErrRecordNotFound
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
	// 判断如果是索引冲突
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		// 断言这里是mysql的错误
		const uniqueConflictsError uint16 = 1062
		if mysqlErr.Number == uniqueConflictsError {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, user domain.User) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", user.Email).First(&u).Error
	return u, err
}

func (dao *UserDAO) UpdateUserInfo(ctx context.Context, user domain.User) error {
	err := dao.db.WithContext(ctx).Save(&user).Error
	return err
}

func (dao *UserDAO) GetById(ctx context.Context, user domain.User) (domain.User, error) {
	var u domain.User
	err := dao.db.WithContext(ctx).Where("id = ?", user.ID).First(&u).Error
	return u, err
}

func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where(`id = ?`, id).First(&u).Error
	return u, err
}

// User 数据库表结构
type User struct {
	Id         int64  `gorm:"primaryKey,autoIncrement" json:"id,omitempty"`
	Email      string `gorm:"unique" json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`

	NickName string `json:"nick_name,omitempty"`
	Birthday string `json:"birthday"`
	Brief    string `json:"brief,omitempty"`
}
