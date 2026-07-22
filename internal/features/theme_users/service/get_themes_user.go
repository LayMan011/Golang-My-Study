package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
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

	lim := 0
	off := 0
	if first != nil {
		lim = *first
	}
	if last != nil {
		off = *last
	}

	cacheKey := fmt.Sprintf("themes_user:limit:%d:offset:%d", lim, off)

	if cached, err := s.themeUserRepositoryRedis.GetThemesUser(ctx, cacheKey); err == nil {
		return cached, nil
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
		return nil, fmt.Errorf("failed save to cache the themes: %w", err)
	}

	return themesUser, nil
}
