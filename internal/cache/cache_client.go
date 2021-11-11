package cache

import (
	"context"

	"github.com/yqchilde/gin-skeleton/internal/model"
	"github.com/yqchilde/gin-skeleton/pkg/cache"
	"github.com/yqchilde/gin-skeleton/pkg/encoding"
	"github.com/yqchilde/gin-skeleton/pkg/redis"
)

func getCacheClient(ctx context.Context) cache.Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	client := cache.NewRedisCache(redis.RC, cachePrefix, jsonEncoding, func() interface{} {
		return &model.User{}
	})
	return client
}
