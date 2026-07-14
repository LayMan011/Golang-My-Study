package core_goredis_pool

import (
	"context"
	"fmt"
	"time"

	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
	opTimeout time.Duration
}

func newRedisClient(config RedisConfig) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.ADDR,
		Password: config.Password,
		DB:       config.Database,
	})

	return &RedisClient{
		rdb,
		config.Timeout,
	}, nil
}

func CreateRedisClientMust(config RedisConfig) *RedisClient {
	client, err := newRedisClient(config)
	if err != nil {
		panic(err)
	}

	return client
}

func (r *RedisClient) Ping(ctx context.Context) core_redis_pool.CustomStatusCmd {
	ans := r.Client.Ping(ctx)
	return &CustomStatusCmd{ans}
}

func (r *RedisClient) GetResult(ctx context.Context, key string) (*string, error) {
	data, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) core_redis_pool.CustomStatusCmd {
	ans := r.Client.Set(ctx, key, value, expiration)
	return &CustomStatusCmd{ans}
}

func (r *RedisClient) HSet(ctx context.Context, key string, value any) error {
	err := r.Client.HSet(ctx, key, value).Err()
	if err != nil {
		return fmt.Errorf("failed to save user to Redis: %w", err)
	}

	return nil
}

func (r *RedisClient) HGetAll(ctx context.Context, login string) error {
	err := r.Client.HGetAll(ctx, login).Err()
	if err != nil {
		return fmt.Errorf("failed to get user to Redis: %w", err)
	}

	return nil
}

func (r *RedisClient) HGetAllResult(ctx context.Context, login string) (map[string]string, error) {
	data, err := r.Client.HGetAll(ctx, login).Result()
	if err != nil {
		return map[string]string{}, fmt.Errorf("failed to get user to Redis: %w", err)
	}

	return data, nil
}

func (r *RedisClient) Expire(ctx context.Context, key string, time time.Duration) error {
	err := r.Client.Expire(ctx, key, time).Err()
	if err != nil {
		return fmt.Errorf("failed to set TTL: %w", err)
	}

	return nil
}

func (r *RedisClient) DoError(ctx context.Context, str1 string, key int) *error {
	ans := r.Client.Do(ctx, str1, key).Err()
	return &ans
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}

func (p *RedisClient) OpTimeout() time.Duration {
	return p.opTimeout
}
