package themes_redis_repository

import (
	"context"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeRepository) DeleteTheme(ctx context.Context, key string) error {
	deleted, err := r.pool.HDel(
		ctx,
		key,
		"id",
		"version",
		"title",
		"description",
		"created_at",
		"subject",
		"rating",
		"all_ratings",
		"number_of_ratings",
		"number_of_users",
		"price",
		"level",
		"duration",
		"format",
		"author_user_id",
	).Result()
	if err != nil {
		return fmt.Errorf("delete theme via HDEL: %w", err)
	}

	if deleted == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
