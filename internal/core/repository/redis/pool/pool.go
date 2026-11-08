package core_redis_pool

import (
	"context"
	"time"
)

type Pool interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) CustomStatusCmd
	Get(ctx context.Context, key string) (string, error)

	HSet(ctx context.Context, key string, value any) error
	HGetAll(ctx context.Context, login string) (map[string]string, error)
	Pipelined(ctx context.Context, fn func(Pipeline) error) (CustomCmder, error)

	HDel(ctx context.Context, key string, fields ...string) CustomIntCmd
	Del(ctx context.Context, keys ...string) error

	Incr(ctx context.Context, key string) (int64, error)
}

type CustomStatusCmd interface {
	Err() error
}

type CustomIntCmd interface {
	Result() (int64, error)
}

type CustomCmder interface {
	Err() error
}

type Pipeline interface {
	HSet(ctx context.Context, key, field string, value any) error
	HDel(ctx context.Context, key string, fields ...string) error
	Expire(ctx context.Context, key string, ttl time.Duration) error
}
