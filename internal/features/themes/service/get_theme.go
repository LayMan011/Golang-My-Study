package themes_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (s *ThemeService) GetTheme(
	ctx context.Context,
	themeID int,
) (domain.Theme, error) {
	key := fmt.Sprintf("theme:%d", themeID)

	theme, err := s.themeRepositoryRedis.GetTheme(ctx, key)
	if err == nil {
		return theme, nil
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.Theme{}, fmt.Errorf("failed to retrieve theme from cache: %w", err)
	}

	theme, err = s.themeRepositoryPostgres.GetTheme(ctx, themeID)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("get theme from repository: %w", err)
	}

	if err := s.themeRepositoryRedis.SaveTheme(ctx, key, theme); err != nil {
		return domain.Theme{}, fmt.Errorf("failed to cache the theme: %w", err)
	}

	return theme, nil
}
