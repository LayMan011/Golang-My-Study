package users_redis_repository

import (
	"context"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) DeleteToken(ctx context.Context, key string) error {
	deleted, err := r.pool.HDel(
		ctx,
		key,
		"access_token",
		"refresh_token",
	).Result()
	if err != nil {
		return fmt.Errorf("delete token via HDEL: %w", err)
	}

	if deleted == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
