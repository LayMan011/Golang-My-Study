package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	lim := -1
	off := -1
	if limit != nil {
		lim = *limit
	}
	if offset != nil {
		off = *offset
	}

	cacheKey := fmt.Sprintf("users:limit:%d:offset:%d", lim, off)

	if cached, err := s.userRepositoryRedis.GetUsers(ctx, cacheKey); err == nil && s.hash_many[cacheKey] {
		return cached, nil
	} else if err != nil {
		s.hash_many[cacheKey] = false
	}

	users, err := s.userRepositoryPostgres.GetUsers(
		ctx,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	if err := s.userRepositoryRedis.SaveUsers(ctx, cacheKey, users); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed save to cache the users", zap.Error(err))
	} else {
		s.hash_many[cacheKey] = true
	}

	return users, nil
}
