package users_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) SaveUserToHash(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	if err := r.pool.DoError(ctx, "SELECT", 0); err != nil {
		return fmt.Errorf("db switch error: %w", err)
	}

	key := fmt.Sprintf("user:%s", user.Login)

	err := r.pool.HSet(ctx, key, user)
	if err != nil {
		return fmt.Errorf("failed to save user to Redis: %w", err)
	}

	err = r.pool.Expire(ctx, key, time.Hour)
	if err != nil {
		return fmt.Errorf("failed to set TTL: %w", err)
	}

	return nil
}
