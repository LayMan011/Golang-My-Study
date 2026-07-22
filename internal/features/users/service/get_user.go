package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	key := fmt.Sprintf("user:%d", id)

	user, err := s.userRepositoryRedis.GetUser(ctx, key)
	if err == nil {
		return user, nil
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.User{}, fmt.Errorf("failed to retrieve user from cache: %w", err)
	}

	user, err = s.userRepositoryPostgres.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := s.userRepositoryRedis.SaveUser(ctx, key, user); err != nil {
		return domain.User{}, fmt.Errorf("failed to cache the user: %w", err)
	}

	return user, nil
}
