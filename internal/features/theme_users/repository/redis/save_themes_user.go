package themes_user_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeUserRepository) SaveThemesUser(ctx context.Context, key string, themesUser []domain.ThemeUser) error {
	b, err := json.Marshal(themesUser)
	if err != nil {
		return fmt.Errorf("marshal themes_user: %w", err)
	}

	if err := r.pool.Set(ctx, key, b, 1*time.Hour).Err(); err != nil {
		return fmt.Errorf("save themes_user to redis: %w", err)
	}

	return nil
}
