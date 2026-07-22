package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err := s.userRepositoryPostgres.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	key := fmt.Sprintf("user:%d", user.ID)

	if err := s.userRepositoryRedis.SaveUser(ctx, key, user); err != nil {
		return domain.User{}, fmt.Errorf("failed to cache the user: %w", err)
	}

	return user, nil
}
