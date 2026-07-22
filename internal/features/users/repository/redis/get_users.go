package users_redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	"github.com/redis/go-redis/v9"
)

func (r *UserRepository) GetUsers(ctx context.Context, key string) ([]domain.User, error) {
	data, err := r.pool.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, core_errors.ErrNotFound
		}
		return nil, fmt.Errorf("get users from redis: %w", err)
	}

	var users []domain.User
	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, fmt.Errorf("unmarshal users from redis: %w", err)
	}

	return users, nil
}
