package themes_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *ThemeRepository) SaveTheme(ctx context.Context, key string, value domain.Theme) error {
	fields := map[string]any{
		"id":                value.ID,
		"version":           value.Version,
		"title":             value.Title,
		"created_at":        value.CreatedAt.Format(time.RFC3339Nano),
		"subject":           value.Subject,
		"all_ratings":       value.AllRatings,
		"number_of_ratings": value.NumberOfRatings,
		"number_of_users":   value.NumberOfUsers,
		"price":             value.Price,
		"level":             value.Level,
		"duration":          value.Duration,
		"format":            value.Format,
		"author_user_id":    value.AuthorUserID,
	}

	if value.Description != nil {
		fields["description"] = *value.Description
	}

	if value.Rating != nil {
		fields["rating"] = *value.Rating
	}

	_, err := r.pool.Pipelined(ctx, func(p core_redis_pool.Pipeline) error {
		for field, fieldValue := range fields {
			if err := p.HSet(ctx, key, field, fieldValue); err != nil {
				return err
			}
		}

		return p.Expire(ctx, key, core_redis_pool.TTLDataSave)
	})

	if err != nil {
		return fmt.Errorf("save theme to Redis: %w", err)
	}

	return nil
}
