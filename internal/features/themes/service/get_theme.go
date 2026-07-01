package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *ThemeService) GetTheme(
	ctx context.Context,
	themeID int,
) (domain.Theme, error) {
	theme, err := s.themeRepository.GetTheme(ctx, themeID)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("get theme from repository: %w", err)
	}

	return theme, nil
}
