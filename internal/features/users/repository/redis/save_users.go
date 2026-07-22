package users_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) SaveUsers(ctx context.Context, key string, users []domain.User) error {
	b, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("marshal users: %w", err)
	}

	if err := r.pool.Set(ctx, key, b, 1*time.Hour).Err(); err != nil {
		return fmt.Errorf("save users to redis: %w", err)
	}

	return nil
}
