package themes_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeService struct {
	themeRepository ThemeRepository
}

type ThemeRepository interface {
	CreateTheme(
		ctx context.Context,
		theme domain.Theme,
	) (domain.Theme, error)

	GetThemes(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Theme, error)

	GetTheme(
		ctx context.Context,
		id int,
	) (domain.Theme, error)

	DeleteTheme(
		ctx context.Context,
		id int,
	) error

	PatchTheme(
		ctx context.Context,
		id int,
		theme domain.Theme,
	) (domain.Theme, error)
}

func NewThemeService(
	themeRepository ThemeRepository,
) *ThemeService {
	return &ThemeService{
		themeRepository: themeRepository,
	}
}
