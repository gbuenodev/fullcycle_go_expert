package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStorage{client: client}, nil
}

func (r *RedisStorage) Increment(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	counterKey := "ratelimit:counter:" + key

	val, err := r.client.Incr(ctx, counterKey).Result()
	if err != nil {
		return 0, err
	}

	if val == 1 {
		r.client.Expire(ctx, counterKey, expiration)
	}

	return val, nil
}

func (r *RedisStorage) IsBlocked(ctx context.Context, key string) (bool, error) {
	blockKey := "ratelimit:block:" + key

	exists, err := r.client.Exists(ctx, blockKey).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (r *RedisStorage) Block(ctx context.Context, key string, duration time.Duration) error {
	blockKey := "ratelimit:block:" + key
	return r.client.Set(ctx, blockKey, "1", duration).Err()
}
