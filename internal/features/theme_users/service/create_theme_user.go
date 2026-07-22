package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
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

	key := fmt.Sprintf("theme_user:%d", themeUser.ID)

	if err := s.themeUserRepositoryRedis.SaveThemeUser(ctx, key, themeUser); err != nil {
		return domain.ThemeUser{}, fmt.Errorf("failed to cache the theme_user: %w", err)
	}

	return themeUserDomain, nil
}
