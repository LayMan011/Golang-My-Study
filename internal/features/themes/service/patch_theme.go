package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *ThemeService) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.ThemePatch,
) (domain.Theme, error) {
	theme, err := s.themeRepositoryPostgres.GetTheme(ctx, id)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("get theme: %w", err)
	}

	if err := theme.ApplyPatch(patch); err != nil {
		return domain.Theme{}, fmt.Errorf("apply theme patch: %w", err)
	}

	patchedTheme, err := s.themeRepositoryPostgres.PatchTheme(ctx, id, theme)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("patch theme: %w", err)
	}

	if err := s.themeRepositoryRedis.PatchTheme(ctx, patchedTheme.ID, patch); err != nil {
		return domain.Theme{}, fmt.Errorf("patch theme to cache: %w", err)
	}

	return patchedTheme, nil
}
