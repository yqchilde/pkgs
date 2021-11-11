package cache

import (
	"context"
	"errors"
	"time"
)

const (
	// DefaultExpireTime Default expiration time
	DefaultExpireTime = time.Hour * 24

	// DefaultNotFoundExpireTime The expiration time when the result is empty is 1 minute
	// which is often used for the cache time when the data is empty (cache penetration)
	DefaultNotFoundExpireTime = time.Minute

	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"
)

var (
	ErrPlaceholder = errors.New("cache: placeholder")
)

// Client Generate a cache client, where keyPrefix is generally the business prefix
var Client Cache

// Cache Define the cache driver interface
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
	Del(ctx context.Context, keys ...string) error
}

// Set data
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, val, expiration)
}

// Get data
func Get(ctx context.Context, key string, val interface{}) error {
	return Client.Get(ctx, key, val)
}

// MultiSet batch set
func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	return Client.MultiSet(ctx, valMap, expiration)
}

// MultiGet batch get
func MultiGet(ctx context.Context, keys []string, valueMap interface{}) error {
	return Client.MultiGet(ctx, keys, valueMap)
}

// Del batch delete
func Del(ctx context.Context, keys ...string) error {
	return Client.Del(ctx, keys...)
}
