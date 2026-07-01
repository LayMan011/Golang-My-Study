package themes_service

import (
	"context"
	"fmt"
)

func (s *ThemeService) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	if err := s.themeRepository.DeleteTheme(ctx, id); err != nil {
		return fmt.Errorf("delete theme from repository: %w", err)
	}

	return nil
}
