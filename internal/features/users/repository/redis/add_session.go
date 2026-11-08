package users_redis_repository

import (
	"context"
	"fmt"
	"time"

	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *UserRepository) AddUserSession(
	ctx context.Context,
	userID int,
	sessionKey string,
	ttl time.Duration,
) error {
	indexKey := fmt.Sprintf("user_sessions:%d", userID)

	cmd, err := r.pool.Pipelined(ctx, func(pipe core_redis_pool.Pipeline) error {
		if err := pipe.HSet(ctx, indexKey, sessionKey, "1"); err != nil {
			return err
		}
		return pipe.Expire(ctx, indexKey, ttl)
	})
	if err != nil {
		return fmt.Errorf("index user session: %w", err)
	}
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("index user session: %w", err)
	}

	return nil
}
