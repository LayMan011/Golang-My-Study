package themes_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeService) CreateTheme(
	ctx context.Context,
	theme domain.Theme,
) (domain.Theme, error) {

	if err := theme.Validate(); err != nil {
		return domain.Theme{}, fmt.Errorf("validate theme domain: %w", err)
	}

	theme, err := s.themeRepositoryPostgres.CreateTheme(ctx, theme)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("create theme: %w", err)
	}

	key := fmt.Sprintf("theme:%d", theme.ID)

	if err := s.themeRepositoryRedis.SaveTheme(ctx, key, theme); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the theme", zap.Error(err))
	}

	return theme, nil
}
