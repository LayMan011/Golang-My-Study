package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *UserService) Login(
	ctx context.Context,
	id int,
	pair *domain.TokenPair,
) error {
	key := fmt.Sprintf("token:%d", id)

	err := s.userRepositoryRedis.SaveToken(ctx, key, pair)
	if err != nil {
		return fmt.Errorf("set token from repository: %w", err)
	}

	return nil
}
