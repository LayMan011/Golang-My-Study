package users_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *UserRepository) SaveUser(ctx context.Context, key string, value domain.User) error {
	fields := map[string]any{
		"id":         value.ID,
		"version":    value.Version,
		"email":      value.Email,
		"created_at": value.CreatedAt.Format(time.RFC3339Nano),
		"password":   string(value.Password),
		"full_name":  value.FullName,
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
		return fmt.Errorf("save user to Redis: %w", err)
	}

	return nil
}
