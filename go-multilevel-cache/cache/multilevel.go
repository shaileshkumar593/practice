package cache

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/singleflight"
)

type MultiLevelCache struct {
	local    *LocalCache
	redis    *RedisCache
	group    singleflight.Group
	localTTL time.Duration
	redisTTL time.Duration
}

func NewMultiLevelCache(
	local *LocalCache,
	redis *RedisCache,
	localTTL, redisTTL time.Duration,
) *MultiLevelCache {
	return &MultiLevelCache{
		local:    local,
		redis:    redis,
		localTTL: localTTL,
		redisTTL: redisTTL,
	}
}

// Get implements cache-aside with:
// L1 local cache -> L2 Redis -> loader/database.
//
// singleflight ensures that concurrent misses for the same key
// execute the loader only once per application instance.
func (c *MultiLevelCache) Get(
	ctx context.Context,
	key string,
	loader func(context.Context) ([]byte, error),
) ([]byte, string, error) {

	// L1
	if value, ok := c.local.Get(key); ok {
		return value, "L1", nil
	}

	// L2
	value, err := c.redis.Get(ctx, key)
	if err == nil {
		c.local.Set(key, value, c.localTTL)
		return value, "L2", nil
	}

	// Coalesce concurrent misses.
	result, err, _ := c.group.Do(key, func() (any, error) {
		// Double-check Redis after entering singleflight.
		value, err := c.redis.Get(ctx, key)
		if err == nil {
			c.local.Set(key, value, c.localTTL)
			return cacheResult{value: value, source: "L2"}, nil
		}

		// Database/source-of-truth loader.
		value, err = loader(ctx)
		if err != nil {
			return nil, err
		}

		// Redis failure should not normally make the request fail.
		// In production, record a metric/log here.
		_ = c.redis.Set(ctx, key, value, c.redisTTL)

		c.local.Set(key, value, c.localTTL)

		return cacheResult{value: value, source: "DB"}, nil
	})

	if err != nil {
		return nil, "", err
	}

	res, ok := result.(cacheResult)
	if !ok {
		return nil, "", errors.New("unexpected cache result type")
	}

	return res.value, res.source, nil
}

type cacheResult struct {
	value  []byte
	source string
}

func (c *MultiLevelCache) Delete(ctx context.Context, key string) error {
	c.local.Delete(key)
	return c.redis.Delete(ctx, key)
}
