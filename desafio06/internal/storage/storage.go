package storage

import (
	"context"
	"time"
)

type Storage interface {
	Increment(ctx context.Context, key string, expiration time.Duration) (int64, error)
	IsBlocked(ctx context.Context, key string) (bool, error)
	Block(ctx context.Context, key string, expiration time.Duration) error
}
