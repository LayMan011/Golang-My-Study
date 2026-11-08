package users_redis_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *UserRepository) SaveToken(ctx context.Context, key string, pair *domain.TokenPair) error {
	if pair == nil {
		return errors.New("save token: nil token pair")
	}

	_, err := r.pool.Pipelined(ctx, func(p core_redis_pool.Pipeline) error {
		if err := p.HSet(ctx, key, "access_token", pair.AccessToken); err != nil {
			return err
		}

		if err := p.HSet(ctx, key, "refresh_token", pair.RefreshToken); err != nil {
			return err
		}

		return p.Expire(ctx, key, time.Hour)
	})

	if err != nil {
		return fmt.Errorf("save token to Redis: %w", err)
	}

	return nil
}
