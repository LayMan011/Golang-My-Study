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
		return fmt.Errorf("failed to delete theme from repository: %w", err)
	}

	if err := s.themeRepositoryRedis.BumpThemesListVersion(ctx); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to invalidate themes list cache", zap.Error(err))
	}

	key := fmt.Sprintf("theme:%d", id)

	if err := s.themeRepositoryRedis.DeleteTheme(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete theme from cache", zap.Error(err))
	}

	return nil
}
