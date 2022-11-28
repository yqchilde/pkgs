package cache

import (
	"errors"
	"strings"
)

// BuildCacheKey Build a cache key with a prefix
func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("cache: key should not be empty")
	}

	cacheKey = key
	if keyPrefix != "" {
		cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	}

	return
}
