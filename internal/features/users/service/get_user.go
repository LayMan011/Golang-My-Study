package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	key := fmt.Sprintf("user:%d", id)

	user, err := s.userRepositoryRedis.GetUser(ctx, key)
	if err == nil && s.hash_once[key] {
		return user, nil
	} else if err != nil {
		s.hash_once[key] = false
		clear(s.hash_many)
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.User{}, fmt.Errorf("failed to retrieve user from cache: %w", err)
	}

	user, err = s.userRepositoryPostgres.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := s.userRepositoryRedis.SaveUser(ctx, key, user); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the user", zap.Error(err))
	} else {
		s.hash_once[key] = true
	}

	return user, nil
}
