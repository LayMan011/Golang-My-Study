package themes_user_service

import (
	"context"
	"fmt"
)

func (s *ThemeUserService) DeleteThemeUser(
	ctx context.Context,
	themeUserID int,
) error {
	if err := s.themeUserRepositoryPostgres.DeleteThemeUser(ctx, themeUserID); err != nil {
		return fmt.Errorf("delete themeUser from repository: %w", err)
	}

	key := fmt.Sprintf("theme_user:%d", themeUserID)

	if err := s.themeUserRepositoryRedis.DeleteThemeUser(ctx, key); err != nil {
		return fmt.Errorf("delete theme_user to cache: %w", err)
	}

	return nil
}
