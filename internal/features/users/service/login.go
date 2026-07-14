package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *UserService) Login(
	ctx context.Context,
	login string,
	pair *domain.TokenPair,
) error {
	err := s.userRepositoryRedis.SaveToken(ctx, login, pair)
	if err != nil {
		return fmt.Errorf("set token from repository: %w", err)
	}

	return nil
}
