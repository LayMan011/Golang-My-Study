package themes_redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	"github.com/redis/go-redis/v9"
)

func (r *ThemeRepository) GetThemes(ctx context.Context, key string) ([]domain.Theme, error) {
	data, err := r.pool.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, core_errors.ErrNotFound
		}
		return nil, fmt.Errorf("get themes from redis: %w", err)
	}

	var themes []domain.Theme
	if err := json.Unmarshal([]byte(data), &themes); err != nil {
		return nil, fmt.Errorf("unmarshal themes from redis: %w", err)
	}

	return themes, nil
}
