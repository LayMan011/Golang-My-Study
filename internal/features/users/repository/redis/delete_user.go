package users_redis_repository

import (
	"context"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, key string) error {
	deleted, err := r.pool.HDel(
		ctx,
		key,
		"id",
		"version",
		"email",
		"created_at",
		"password",
		"full_name",
	).Result()
	if err != nil {
		return fmt.Errorf("delete user via HDEL: %w", err)
	}

	if deleted == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
