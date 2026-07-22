package themes_service

import (
	"context"
	"fmt"
)

func (s *ThemeService) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	if err := s.themeRepositoryPostgres.DeleteTheme(ctx, id); err != nil {
		return fmt.Errorf("delete theme from repository: %w", err)
	}

	key := fmt.Sprintf("theme:%d", id)

	if err := s.themeRepositoryRedis.DeleteTheme(ctx, key); err != nil {
		return fmt.Errorf("delete theme to cache: %w", err)
	}

	return nil
}
