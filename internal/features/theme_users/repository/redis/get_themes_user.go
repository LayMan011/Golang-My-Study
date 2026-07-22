package themes_user_redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	"github.com/redis/go-redis/v9"
)

func (r *ThemeUserRepository) GetThemesUser(ctx context.Context, key string) ([]domain.ThemeUser, error) {
	data, err := r.pool.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, core_errors.ErrNotFound
		}
		return nil, fmt.Errorf("get themes_user from redis: %w", err)
	}

	var themesUser []domain.ThemeUser
	if err := json.Unmarshal([]byte(data), &themesUser); err != nil {
		return nil, fmt.Errorf("unmarshal themes_user from redis: %w", err)
	}

	return themesUser, nil
}
