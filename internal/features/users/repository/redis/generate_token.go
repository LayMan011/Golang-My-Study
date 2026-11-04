package users_redis_repository

import (
	"context"
	"os"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	accessTTL  = 15 * time.Minute
	refreshTTL = 7 * 24 * time.Hour
)

func (r *UserRepository) GenerateTokenPair(ctx context.Context, ID string) (*domain.TokenPair, error) {
	secret := []byte(os.Getenv("SECTER_KEY_FOR_JWT"))

	accessClaims := domain.NewClaims(ID, uuid.NewString(), accessTTL)
	refreshClaims := domain.NewClaims(ID, uuid.NewString(), refreshTTL)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secret)
	if err != nil {
		return nil, err
	}

	if err := r.pool.Set(ctx, "refresh:"+ID, ID, refreshTTL); err.Err() != nil {
		return nil, err.Err()
	}

	return domain.NewTokenPair(accessToken, refreshToken), nil
}
