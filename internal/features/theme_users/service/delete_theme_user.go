package themes_user_service

import (
	"context"
	"fmt"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeUserService) DeleteThemeUser(
	ctx context.Context,
	themeUserID int,
) error {
	if err := s.themeUserRepositoryPostgres.DeleteThemeUser(ctx, themeUserID); err != nil {
		return fmt.Errorf("failed to delete theme_user from repository: %w", err)
	}

	key := fmt.Sprintf("theme_user:%d", themeUserID)
	s.hash_once[key] = false
	clear(s.hash_many)

	if err := s.themeUserRepositoryRedis.DeleteThemeUser(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete theme_user from cache", zap.Error(err))
	} else {
		s.hash_once[key] = false
		clear(s.hash_many)
	}

	return nil
}
