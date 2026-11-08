package themes_user_redis_repository

import (
	"context"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeUserRepository) DeleteThemeUser(ctx context.Context, key string) error {
	deleted, err := r.pool.HDel(
		ctx,
		key,
		"id",
		"version",
		"completed",
		"addition_at",
		"completed_at",
		"percentages",
		"total_lessons",
		"completed_lessons",
		"theme_id",
		"user_id",
	).Result()
	if err != nil {
		return fmt.Errorf("delete ThemeUser via HDEL: %w", err)
	}

	if deleted == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
