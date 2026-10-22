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
	clear(s.hash_many)

	if err := s.userRepositoryRedis.SaveUser(ctx, key, user); err != nil {
<<<<<<< HEAD
		return domain.User{}, fmt.Errorf("failed to cache the user: %w", err)
=======
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the user", zap.Error(err))
	} else {
		s.hash_once[key] = true
>>>>>>> c8567b9 (fixed cache and add tests)
	}

	return user, nil
}
