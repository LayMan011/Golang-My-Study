package users_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *UserRepository) SaveToken(
	ctx context.Context,
	key string,
	pair *domain.TokenPair,
) error {
	if pair == nil {
		return fmt.Errorf("token pair is nil")
	}

	accessClaims, err := domain.ParseToken(pair.AccessToken)
	if err != nil {
		return fmt.Errorf("parse access token: %w", err)
	}

	refreshClaims, err := domain.ParseToken(pair.RefreshToken)
	if err != nil {
		return fmt.Errorf("parse refresh token: %w", err)
	}

	if accessClaims.ID == "" || refreshClaims.ID == "" {
		return fmt.Errorf("token is missing jti")
	}

	accessKey := "token:" + accessClaims.ID
	if key != accessKey {
		return fmt.Errorf("access token key %q does not match expected key %q", key, accessKey)
	}

	if refreshClaims.ExpiresAt == nil {
		return fmt.Errorf("refresh token is missing expiration")
	}

	refreshTTL := time.Until(refreshClaims.ExpiresAt.Time)
	if refreshTTL <= 0 {
		return fmt.Errorf("refresh token is expired")
	}

	const tokenTTL = time.Hour
	indexKey := fmt.Sprintf("user_sessions:%d", accessClaims.UserID)

	indexTTL := refreshTTL
	if tokenTTL > indexTTL {
		indexTTL = tokenTTL
	}

	cmd, err := r.pool.Pipelined(ctx, func(p core_redis_pool.Pipeline) error {
		if err := p.HSet(ctx, key, "access_token", pair.AccessToken); err != nil {
			return err
		}
		if err := p.HSet(ctx, key, "refresh_token", pair.RefreshToken); err != nil {
			return err
		}
		if err := p.Expire(ctx, key, tokenTTL); err != nil {
			return err
		}

<<<<<<< HEAD
		if err := p.Expire(ctx, key, tokenTTL); err != nil {
			return err
		}

=======
>>>>>>> cfdc737 (mini fix pull)
		if err := p.HSet(ctx, indexKey, accessKey, "1"); err != nil {
			return err
		}
		if err := p.HSet(
			ctx,
			indexKey,
			"refresh:"+refreshClaims.ID,
			"1",
		); err != nil {
			return err
		}
		return p.Expire(ctx, indexKey, indexTTL)
	})
	if err != nil {
		return fmt.Errorf("save token and index user sessions: %w", err)
	}
	
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("save token and index user sessions: %w", err)
	}

	return nil
}
