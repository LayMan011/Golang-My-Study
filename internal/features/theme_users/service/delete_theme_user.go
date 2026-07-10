package themes_user_service

import (
	"context"
	"fmt"
)

func (s *ThemeUserService) DeleteThemeUser(
	ctx context.Context,
	themeUserID int,
) error {
	if err := s.themeUserRepository.DeleteThemeUser(ctx, themeUserID); err != nil {
		return fmt.Errorf("delete themeUser from repository: %w", err)
	}

	return nil
}
