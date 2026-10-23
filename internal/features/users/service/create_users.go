package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
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
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the user", zap.Error(err))
	} else {
		s.hash_once[key] = true
	}

	return user, nil
}
