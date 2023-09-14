package cache

import "github.com/redis/go-redis/v9"

type UserCache struct {
	client redis.Cmdable
}

// NewUserCache 依赖注入，绝对不内部初始化，要通过外面注入进来
// 1. 用接口 2. 定义成字段 3. 外面传进来
func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client: client,
	}
}
