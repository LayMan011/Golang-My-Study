package users_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) SaveToken(ctx context.Context, login string, pair *domain.TokenPair) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	if err := r.pool.DoError(ctx, "SELECT", 0); err != nil {
		return fmt.Errorf("db switch error: %w", err)
	}

	key := fmt.Sprintf("jwt:%s", login)

	jsonData, err := json.Marshal(pair)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	err = r.pool.Set(ctx, key, jsonData, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("db set error: %w", err)
	}

	return nil
}
