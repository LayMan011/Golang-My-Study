package core_redis_pool

import (
	"context"
	"time"
)

type Pool interface {
	// Ping(ctx context.Context) CustomStatusCmd

	Set(ctx context.Context, key string, value any, expiration time.Duration) CustomStatusCmd // +
	Get(ctx context.Context, key string) (string, error)                                      // +

	DoError(ctx context.Context, str1 string, key int) error
	HSet(ctx context.Context, key string, value any) error                       // +
	HGetAll(ctx context.Context, login string) (map[string]string, error)        // +
	Expire(ctx context.Context, key string, time time.Duration) error            // +
	Pipelined(ctx context.Context, fn func(Pipeline) error) (CustomCmder, error) // +

	Exists(ctx context.Context, key string) CustomIntCmd                 // +
	HDel(ctx context.Context, key string, fields ...string) CustomIntCmd // +

	Close() error // +
	Incr(ctx context.Context, key string) (int64, error)

	// OpTimeout() time.Duration
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
	Expire(ctx context.Context, key string, ttl time.Duration) error
}
