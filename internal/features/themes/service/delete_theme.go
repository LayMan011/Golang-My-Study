package themes_service

import (
	"context"
	"fmt"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeService) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	if err := s.themeRepositoryPostgres.DeleteTheme(ctx, id); err != nil {
		return fmt.Errorf("delete theme from repository: %w", err)
	}

	theme_user_ids, err := s.themeUserCachePostgres.GetIDsByThemeID(ctx, id)
	if err != nil {
		return fmt.Errorf("get theme_user IDs for theme %d: %w", id, err)
	}

	if err := s.themeUserCacheRedis.InvalidateByIDs(ctx, theme_user_ids); err != nil {
		return fmt.Errorf("invalidate theme_user cache for theme %d: %w", id, err)
	}

	if err := s.themeRepositoryRedis.BumpThemesListVersion(ctx); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to get user IDs for theme cache invalidation", zap.Int("theme_id", id), zap.Error(err))
	}

	key := fmt.Sprintf("theme:%d", id)

	if err := s.themeRepositoryRedis.DeleteTheme(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete theme from cache", zap.Error(err))
	}

	return nil
}
