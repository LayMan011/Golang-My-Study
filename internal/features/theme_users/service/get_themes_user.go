package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeUserService) GetThemesUser(
	ctx context.Context,
	userID *int,
	themeID *int,
	first *int,
	last *int,
) ([]domain.ThemeUser, error) {
	if first != nil && *first < 0 {
		return nil, fmt.Errorf(
			"'first' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if last != nil && *last < 0 {
		return nil, fmt.Errorf(
			"'last' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	lim := -1
	off := -1
	if first != nil {
		lim = *first
	}
	if last != nil {
		off = *last
	}

	cacheKey := fmt.Sprintf("themes_user:limit:%d:offset:%d", lim, off)

	if cached, err := s.themeUserRepositoryRedis.GetThemesUser(ctx, cacheKey); err == nil && s.hash_many[cacheKey] {
		return cached, nil
	} else if err != nil {
		s.hash_many[cacheKey] = false
	}

	themesUser, err := s.themeUserRepositoryPostgres.GetThemesUser(
		ctx,
		userID,
		themeID,
		first,
		last,
	)
	if err != nil {
		return nil, fmt.Errorf("get themes_user from repository: %w", err)
	}

	if err := s.themeUserRepositoryRedis.SaveThemesUser(ctx, cacheKey, themesUser); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed save to cache the themes", zap.Error(err))
	} else {
		s.hash_many[cacheKey] = true
	}

	return themesUser, nil
}
