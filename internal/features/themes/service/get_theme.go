package themes_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeService) GetTheme(
	ctx context.Context,
	themeID int,
) (domain.Theme, error) {
	key := fmt.Sprintf("theme:%d", themeID)

	theme, err := s.themeRepositoryRedis.GetTheme(ctx, key)
	if err == nil && s.hash_once[key] {
		return theme, nil
	} else if err != nil {
		s.hash_once[key] = false
		clear(s.hash_many)
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.Theme{}, fmt.Errorf("failed to retrieve theme from cache: %w", err)
	}

	theme, err = s.themeRepositoryPostgres.GetTheme(ctx, themeID)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("get theme from repository: %w", err)
	}

	if err := s.themeRepositoryRedis.SaveTheme(ctx, key, theme); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the theme", zap.Error(err))
	} else {
		s.hash_once[key] = true
	}

	return theme, nil
}
