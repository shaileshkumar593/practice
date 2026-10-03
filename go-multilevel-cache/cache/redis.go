package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 10,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})

	return &RedisCache{client: client}
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	return r.client.Get(ctx, key).Bytes()
}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.client.Ping(ctx).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
