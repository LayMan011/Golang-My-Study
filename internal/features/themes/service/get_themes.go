package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeService) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"'limit' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"'offset' must be non-negative: %w",
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

	cacheKey := fmt.Sprintf("themes:limit:%d:offset:%d", lim, off)

	if cached, err := s.themeRepositoryRedis.GetThemes(ctx, cacheKey); err == nil && s.hash_many[cacheKey] {
		return cached, nil
	}

	themes, err := s.themeRepositoryPostgres.GetThemes(
		ctx,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get themes from repository: %w", err)
	}

	if err := s.themeRepositoryRedis.SaveThemes(ctx, cacheKey, themes); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed save to cache the themes", zap.Error(err))
	} else {
		s.hash_many[cacheKey] = true
	}

	return themes, nil
}
