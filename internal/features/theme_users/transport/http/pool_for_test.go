package themes_user_transport_http

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type getThemesUserServiceStub struct {
	themes    []domain.ThemeUser
	gotLimit  *int
	gotOffset *int
}

func (s *getThemesUserServiceStub) GetThemesUser(
	_ context.Context,
	_ *int,
	_ *int,
	limit *int,
	offset *int,
) ([]domain.ThemeUser, error) {
	s.gotLimit = limit
	s.gotOffset = offset

	start, end := 0, len(s.themes)
	if offset != nil {
		start = *offset
	}

	if limit != nil {
		end = start + *limit
	}

	if start >= len(s.themes) {
		return []domain.ThemeUser{}, nil
	}

	if end > len(s.themes) {
		end = len(s.themes)
	}

	return s.themes[start:end], nil
}

func (s *getThemesUserServiceStub) CreateThemeUser(
	ctx context.Context,
	themeUser domain.ThemeUser,
) (domain.ThemeUser, error) {
	return domain.ThemeUser{}, nil
}

func (s *getThemesUserServiceStub) GetThemeUser(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	return domain.ThemeUser{}, nil
}

func (s *getThemesUserServiceStub) DeleteThemeUser(
	ctx context.Context,
	themeUserID int,
) error {
	return nil
}

func (s *getThemesUserServiceStub) PatchThemeUser(
	ctx context.Context,
	themeUserID int,
	themeUserPatch domain.ThemeUserPatch,
) (domain.ThemeUser, error) {
	return domain.ThemeUser{}, nil
}
