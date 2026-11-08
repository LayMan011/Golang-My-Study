package themes_user_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *ThemeUserRepository) SaveThemeUser(ctx context.Context, key string, value domain.ThemeUser) error {
	fields := map[string]any{
		"id":                value.ID,
		"version":           value.Version,
		"completed":         value.Completed,
		"addition_at":       value.AdditionAt.Format(time.RFC3339Nano),
		"percentages":       value.Percentages,
		"total_lessons":     value.TotalLessons,
		"completed_lessons": value.CompletedLessons,
		"theme_id":          value.ThemeID,
		"user_id":           value.UserID,
	}

	if value.CompletedAt != nil {
		fields["completed_at"] = value.CompletedAt.Format(time.RFC3339Nano)
	}

	_, err := r.pool.Pipelined(ctx, func(p core_redis_pool.Pipeline) error {
		if value.CompletedAt == nil {
			if err := p.HDel(ctx, key, "completed_at"); err != nil {
				return err
			}
		}

		for field, fieldValue := range fields {
			if err := p.HSet(ctx, key, field, fieldValue); err != nil {
				return err
			}
		}

		return p.Expire(ctx, key, core_redis_pool.TTLDataSave)
	})

	if err != nil {
		return fmt.Errorf("save theme_user to Redis: %w", err)
	}

	return nil
}
