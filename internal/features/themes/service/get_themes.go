package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (s *ThemeService) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
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

	themes, err := s.themeRepository.GetThemes(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get themes from repository: %w", err)
	}

	return themes, nil
}
