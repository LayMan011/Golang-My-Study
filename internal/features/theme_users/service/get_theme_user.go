package themes_user_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *ThemeUserService) GetThemeUser(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	themeUser, err := s.themeUserRepository.GetThemeUser(ctx, id)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("get theme_user from repository: %w", err)
	}

	return themeUser, nil
}
