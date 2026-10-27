package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeUserService) PatchThemeUser(
	ctx context.Context,
	themeUserID int,
	themeUserPatch domain.ThemeUserPatch,
) (domain.ThemeUser, error) {
	theme, err := s.GetThemeUser(ctx, themeUserID)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("get theme: %w", err)
	}

	if err := theme.ApplyPatch(themeUserPatch); err != nil {
		return domain.ThemeUser{}, fmt.Errorf("apply theme patch: %w", err)
	}

	patchedTheme, err := s.themeUserRepositoryPostgres.PatchThemeUser(ctx, themeUserID, theme)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("patch theme: %w", err)
	}

	if err := s.themeUserRepositoryRedis.BumpThemesUserListVersion(ctx); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to invalidate themes_user list cache", zap.Error(err))
	}

	if err := s.themeUserRepositoryRedis.PatchThemeUser(ctx, patchedTheme.ID, themeUserPatch); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("patch theme to cache", zap.Error(err))
	}

	return patchedTheme, nil
}
