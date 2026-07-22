package core_goredis_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CustomStringCmd struct {
	*redis.StringCmd
}

type CustomStatusCmd struct {
	*redis.StatusCmd
}

type CustomCmder struct {
	cmds []redis.Cmder
}

type goRedisPipeline struct {
	pipe redis.Pipeliner
}

type CustomIntCmd struct {
	*redis.IntCmd
}

func (c CustomIntCmd) Result() (int64, error) {
	if c.IntCmd == nil {
		return 0, fmt.Errorf("nil IntCmd")
	}
	return c.IntCmd.Result()
}

func (cmd *CustomStringCmd) Result() (string, error) {
	return cmd.StringCmd.Result()
}

func (status *CustomStatusCmd) Err() error {
	return status.StatusCmd.Err()
}

func (c *CustomCmder) Err() error {
	if c == nil || len(c.cmds) == 0 {
		return nil
	}

	for _, cmd := range c.cmds {
		if err := cmd.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (p *goRedisPipeline) Expire(ctx context.Context, key string, ttl time.Duration) error {
	cmd := p.pipe.Expire(ctx, key, ttl)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("failed to set TTL: %w", err)
	}
	return nil
}

func (p *goRedisPipeline) HSet(ctx context.Context, key string, field string, value any) error {
	cmd := p.pipe.HSet(ctx, key, field, value)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("failed to save to Redis hash: %w", err)
	}
	return nil
}
