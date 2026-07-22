package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
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

	lim := 0
	off := 0
	if limit != nil {
		lim = *limit
	}
	if offset != nil {
		off = *offset
	}

	cacheKey := fmt.Sprintf("users:limit:%d:offset:%d", lim, off)

	if cached, err := s.userRepositoryRedis.GetUsers(ctx, cacheKey); err == nil {
		return cached, nil
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
		return nil, fmt.Errorf("failed save to cache the users: %w", err)
	}

	return users, nil
}
