package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
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

	key := fmt.Sprintf("theme:%d", patchedTheme.ID)
	clear(s.hash_many)

	if err := s.themeRepositoryRedis.PatchTheme(ctx, patchedTheme.ID, patch); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("patch theme to cache", zap.Error(err))
	} else {
		s.hash_once[key] = false
	}

	return patchedTheme, nil
}
