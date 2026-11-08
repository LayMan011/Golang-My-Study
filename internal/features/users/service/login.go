package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *UserService) Login(
	ctx context.Context,
	pair *domain.TokenPair,
) error {
	claims, err := domain.ParseToken(pair.AccessToken)
	if err != nil {
		return fmt.Errorf("parse access token: %w", err)
	}

	if claims.ID == "" {
		return fmt.Errorf("access token is missing jti")
	}

	key := "token:" + claims.ID
	if err := s.userRepositoryRedis.SaveToken(ctx, key, pair); err != nil {
		return fmt.Errorf("save token to repository: %w", err)
	}

	return nil
}
