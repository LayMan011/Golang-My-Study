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

	themesUser, err := s.themeUserRepository.GetThemesUser(ctx, userID, themeID, first, last)
	if err != nil {
		return []domain.ThemeUser{}, fmt.Errorf("get themes_user from repository: %w", err)
	}

	return themesUser, nil
}
