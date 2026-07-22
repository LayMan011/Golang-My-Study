package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserService) GenerateTokenPair(
	ctx context.Context,
	login string,
) (*domain.TokenPair, error) {
	token, err := r.userRepositoryRedis.GenerateTokenPair(ctx, login)
	if err != nil {
		return &domain.TokenPair{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
