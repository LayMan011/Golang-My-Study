package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *ThemeUserService) CreateThemeUser(
	ctx context.Context,
	themeUser domain.ThemeUser,
) (domain.ThemeUser, error) {
	if err := themeUser.Validate(); err != nil {
		return domain.ThemeUser{}, fmt.Errorf("validate themeUser domain: %w", err)
	}

	themeUserDomain, err := s.themeUserRepositoryPostgres.CreateThemeUser(ctx, themeUser)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("create themeUser: %w", err)
	}

	key := fmt.Sprintf("theme_user:%d", themeUserDomain.ID)
	clear(s.hash_many)

	if err := s.themeUserRepositoryRedis.SaveThemeUser(ctx, key, themeUserDomain); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to cache the theme_user", zap.Error(err))
	} else {
		s.hash_once[key] = true
	}

	return themeUserDomain, nil
}
