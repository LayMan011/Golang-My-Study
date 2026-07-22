package themes_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeRepository) SaveThemes(ctx context.Context, key string, themes []domain.Theme) error {
	b, err := json.Marshal(themes)
	if err != nil {
		return fmt.Errorf("marshal themes: %w", err)
	}

	if err := r.pool.Set(ctx, key, b, 1*time.Hour).Err(); err != nil {
		return fmt.Errorf("save themes to redis: %w", err)
	}

	return nil
}
