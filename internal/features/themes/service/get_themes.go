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
	result, err := getQueryParams(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	version, err := s.themeRepositoryRedis.GetThemesListVersion(ctx)
	if err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to get themes list version", zap.Error(err))
	}

	cacheKey := fmt.Sprintf("themes:list:version:%d:user_id:%d:limit:%d:offset:%d", version, result[0], result[1], result[2])

	if cached, err := s.themeRepositoryRedis.GetThemes(ctx, cacheKey); err == nil {
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
	}

	return themes, nil
}

func getQueryParams(userID *int, limit *int, offset *int) ([3]int, error) {
	if limit != nil && *limit < 0 {
		return [3]int{}, fmt.Errorf(
			"'first' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return [3]int{}, fmt.Errorf(
			"'last' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if userID != nil && *userID < 0 {
		return [3]int{}, fmt.Errorf(
			"'user_id' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	lim := -1
	off := -1
	user_id := -1
	if limit != nil {
		lim = *limit
	}
	if offset != nil {
		off = *offset
	}
	if userID != nil {
		user_id = *userID
	}

	return [3]int{user_id, lim, off}, nil
}
