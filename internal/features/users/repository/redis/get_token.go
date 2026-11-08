package users_redis_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) GetToken(ctx context.Context, key string) (domain.TokenPair, error) {
	data, err := r.pool.HGetAll(ctx, key)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("get token from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.TokenPair{}, core_errors.ErrNotFound
	}

	return domain.TokenPair{
		AccessToken:  data["access_token"],
		RefreshToken: data["refresh_token"],
	}, nil
}
