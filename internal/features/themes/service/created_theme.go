package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *ThemeService) CreateTheme(
	ctx context.Context,
	theme domain.Theme,
) (domain.Theme, error) {

	if err := theme.Validate(); err != nil {
		return domain.Theme{}, fmt.Errorf("validate theme domain: %w", err)
	}

	theme, err := s.themeRepository.CreateTheme(ctx, theme)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("create theme: %w", err)
	}

	return theme, nil
}
