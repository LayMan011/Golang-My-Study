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
	limit *int,
	offset *int,
) ([]domain.ThemeUser, error) {
	result, err := getQueryParams(userID, themeID, limit, offset)
	if err != nil {
		return nil, err
	}

	version, err := s.themeUserRepositoryRedis.GetThemesUserListVersion(ctx)
	if err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to get themes_user list version", zap.Error(err))
	}

	cacheKey := fmt.Sprintf("themes_user:list:version:%duser_id:%dtheme_id:%dlimit:%d:offset:%d", version, result[0], result[1], result[2], result[3])

	if cached, err := s.themeUserRepositoryRedis.GetThemesUser(ctx, cacheKey); err == nil {
		return cached, nil
	}

	themesUser, err := s.themeUserRepositoryPostgres.GetThemesUser(
		ctx,
		userID,
		themeID,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get themes_user from repository: %w", err)
	}

	if err := s.themeUserRepositoryRedis.SaveThemesUser(ctx, cacheKey, themesUser); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed save to cache the themes", zap.Error(err))
	}

	return themesUser, nil
}

func getQueryParams(userID *int, themeID *int, limit *int, offset *int) ([4]int, error) {
	if limit != nil && *limit < 0 {
		return [4]int{}, fmt.Errorf(
			"'first' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return [4]int{}, fmt.Errorf(
			"'last' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if userID != nil && *userID < 0 {
		return [4]int{}, fmt.Errorf(
			"'user_id' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if themeID != nil && *themeID < 0 {
		return [4]int{}, fmt.Errorf(
			"'theme_id' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	lim := -1
	off := -1
	user_id := -1
	theme_id := -1
	if limit != nil {
		lim = *limit
	}
	if offset != nil {
		off = *offset
	}
	if userID != nil {
		user_id = *userID
	}
	if themeID != nil {
		theme_id = *themeID
	}

	return [4]int{user_id, theme_id, lim, off}, nil
}
