package themes_user_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (s *ThemeUserService) GetThemeUser(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	key := fmt.Sprintf("theme_user:%d", id)

	themeUser, err := s.themeUserRepositoryRedis.GetThemeUser(ctx, key)
	if err == nil {
		return themeUser, nil
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.ThemeUser{}, fmt.Errorf("failed to retrieve theme_user from cache: %w", err)
	}

	themeUser, err = s.themeUserRepositoryPostgres.GetThemeUser(ctx, id)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("get theme_user from repository: %w", err)
	}

	if err := s.themeUserRepositoryRedis.SaveThemeUser(ctx, key, themeUser); err != nil {
		return domain.ThemeUser{}, fmt.Errorf("failed to cache the theme_user: %w", err)
	}

	return themeUser, nil
}
