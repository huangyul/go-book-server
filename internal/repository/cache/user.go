package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-book-server/internal/domain"
	"time"
)

type UserCache struct {
	client     redis.Cmdable
	expiration time.Duration
}

// NewUserCache 依赖注入，绝对不内部初始化，要通过外面注入进来
// 1. 用接口 2. 定义成字段 3. 外面传进来
func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

func (c *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := c.key(id)
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, nil
	}
	var u domain.User
	err = json.Unmarshal(val, &u)
	return u, err
}

func (c *UserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := c.key(u.ID)
	return c.client.Set(ctx, key, val, c.expiration).Err()
}

func (c *UserCache) key(id int64) string {
	return fmt.Sprintf(`user:info:%d`, id)
}
