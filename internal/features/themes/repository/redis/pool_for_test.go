package themes_redis_repository

import (
	"context"
	"time"

	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
	"github.com/redis/go-redis/v9"
)

type testRedisPool struct {
	client *redis.Client
}

var _ core_redis_pool.Pool = (*testRedisPool)(nil)

func (p *testRedisPool) Set(
	ctx context.Context,
	key string,
	value any,
	expiration time.Duration,
) core_redis_pool.CustomStatusCmd {
	return p.client.Set(ctx, key, value, expiration)
}

func (p *testRedisPool) Get(ctx context.Context, key string) (string, error) {
	return p.client.Get(ctx, key).Result()
}

func (p *testRedisPool) DoError(ctx context.Context, str1 string, key int) error {
	return p.client.Do(ctx, str1, key).Err()
}

func (p *testRedisPool) HSet(ctx context.Context, key string, value any) error {
	return p.client.HSet(ctx, key, value).Err()
}

func (p *testRedisPool) HGetAll(
	ctx context.Context,
	key string,
) (map[string]string, error) {
	return p.client.HGetAll(ctx, key).Result()
}

func (p *testRedisPool) Expire(
	ctx context.Context,
	key string,
	expiration time.Duration,
) error {
	return p.client.Expire(ctx, key, expiration).Err()
}

func (p *testRedisPool) Pipelined(
	ctx context.Context,
	fn func(core_redis_pool.Pipeline) error,
) (core_redis_pool.CustomCmder, error) {
	cmds, err := p.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		return fn(testRedisPipeline{pipe: pipe})
	})

	return testRedisPipelineResult{
		cmds: cmds,
		err:  err,
	}, err
}

func (p *testRedisPool) Exists(
	ctx context.Context,
	key string,
) core_redis_pool.CustomIntCmd {
	return p.client.Exists(ctx, key)
}

func (p *testRedisPool) HDel(
	ctx context.Context,
	key string,
	fields ...string,
) core_redis_pool.CustomIntCmd {
	return p.client.HDel(ctx, key, fields...)
}

func (p *testRedisPool) Close() error {
	return p.client.Close()
}

type testRedisPipeline struct {
	pipe redis.Pipeliner
}

func (p testRedisPipeline) HSet(
	ctx context.Context,
	key string,
	field string,
	value any,
) error {
	p.pipe.HSet(ctx, key, field, value)
	return nil
}

func (p testRedisPipeline) Expire(
	ctx context.Context,
	key string,
	expiration time.Duration,
) error {
	p.pipe.Expire(ctx, key, expiration)
	return nil
}

type testRedisPipelineResult struct {
	cmds []redis.Cmder
	err  error
}

func (r testRedisPipelineResult) Err() error {
	if r.err != nil {
		return r.err
	}

	for _, cmd := range r.cmds {
		if err := cmd.Err(); err != nil {
			return err
		}
	}

	return nil
}
