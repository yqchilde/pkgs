package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/yqchilde/gin-skeleton/internal/model"
	"github.com/yqchilde/gin-skeleton/pkg/cache"
	"github.com/yqchilde/gin-skeleton/pkg/encoding"
	"github.com/yqchilde/gin-skeleton/pkg/redis"
)

const (
	PrefixUserBaseCacheKey = "user:base:%s"
)

type UserBaseCache struct {
	cache cache.Cache
}

func NewUserBaseCache() *UserBaseCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	return &UserBaseCache{
		cache: cache.NewRedisCache(redis.RC, cachePrefix, jsonEncoding, func() interface{} {
			return &model.User{}
		}),
	}
}

// GetUserBaseCacheKey 获取cache key
func (c *UserBaseCache) GetUserBaseCacheKey(userID string) string {
	return fmt.Sprintf(PrefixUserBaseCacheKey, userID)
}

// SetUserBaseCache Write to user cache
func (c *UserBaseCache) SetUserBaseCache(ctx context.Context, userID string, user *model.User, duration time.Duration) error {
	if user == nil {
		return nil
	}
	cacheKey := c.GetUserBaseCacheKey(userID)
	err := c.cache.Set(ctx, cacheKey, user, duration)
	if err != nil {
		return errors.Wrap(err, "[cache.user_base] set user base cache err")
	}
	return nil
}

// GetUserBaseCache Get user cache
func (c *UserBaseCache) GetUserBaseCache(ctx context.Context, userID string) (data *model.User, err error) {
	cacheKey := c.GetUserBaseCacheKey(userID)
	err = c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[cache.user_base] set user base cache err")
	}
	return data, nil
}
