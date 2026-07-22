package core_goredis_pool

import (
	"context"
	"errors"
	"fmt"
	"time"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
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

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

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
	return &CustomStatusCmd{r.Client.Ping(ctx)}
}

func (r *RedisClient) GetResult(ctx context.Context, key string) (*string, error) {
	data, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get key %s: %w", key, err)
	}
	return &data, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) core_redis_pool.CustomStatusCmd {
	return &CustomStatusCmd{r.Client.Set(ctx, key, value, expiration)}
}

func (r *RedisClient) HSet(ctx context.Context, key string, value any) error {
	if err := r.Client.HSet(ctx, key, value).Err(); err != nil {
		return fmt.Errorf("failed to save to Redis hash: %w", err)
	}
	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", core_errors.ErrNotFound
		}
		return "", fmt.Errorf("failed to get value from redis: %w", err)
	}

	return val, nil
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	data, err := r.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get hash from Redis: %w", err)
	}
	return data, nil
}

func (r *RedisClient) HGetAllResult(ctx context.Context, login string) (map[string]string, error) {
	data, err := r.Client.HGetAll(ctx, login).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user to Redis: %w", err)
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

func (r *RedisClient) Pipelined(
	ctx context.Context,
	fn func(core_redis_pool.Pipeline) error,
) (core_redis_pool.CustomCmder, error) {
	cmds, err := r.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		return fn(&goRedisPipeline{pipe: pipe})
	})
	if err != nil {
		return nil, err
	}

	return &CustomCmder{cmds: cmds}, nil
}

func (r *RedisClient) Exists(ctx context.Context, key string) core_redis_pool.CustomIntCmd {
	return CustomIntCmd{r.Client.Exists(ctx, key)}
}

func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) core_redis_pool.CustomIntCmd {
	return CustomIntCmd{r.Client.HDel(ctx, key, fields...)}
}

func (r *RedisClient) DoError(ctx context.Context, str1 string, key int) error {
	return r.Client.Do(ctx, str1, key).Err()
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}

func (p *RedisClient) OpTimeout() time.Duration {
	return p.opTimeout
}
