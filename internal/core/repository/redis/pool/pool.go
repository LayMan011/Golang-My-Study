package core_redis_pool

import (
	"context"
	"time"
)

type Pool interface {
	Ping(ctx context.Context) CustomStatusCmd
	GetResult(ctx context.Context, key string) (*string, error)

	Set(ctx context.Context, key string, value any, expiration time.Duration) CustomStatusCmd
	Get(ctx context.Context, key string) (string, error)

	DoError(ctx context.Context, str1 string, key int) error
	HSet(ctx context.Context, key string, value any) error
	HGetAll(ctx context.Context, login string) (map[string]string, error)
	HGetAllResult(ctx context.Context, login string) (map[string]string, error)
	Expire(ctx context.Context, key string, time time.Duration) error
	Pipelined(ctx context.Context, fn func(Pipeline) error) (CustomCmder, error)

	Exists(ctx context.Context, key string) CustomIntCmd
	HDel(ctx context.Context, key string, fields ...string) CustomIntCmd

	Close() error

	OpTimeout() time.Duration
}

type CustomStatusCmd interface {
	Err() error
}

type CustomStringCmd interface {
	Result() (string, error)
}

type CustomCmd interface {
	Result() string
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
