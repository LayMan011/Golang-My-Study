package themes_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeService struct {
	themeRepositoryPostgres ThemeRepositoryPostgres
	themeRepositoryRedis    ThemeRepositoryRedis
}

type ThemeRepositoryRedis interface {
	SaveTheme(
		ctx context.Context,
		key string,
		value domain.Theme,
	) error

	GetTheme(
		ctx context.Context,
		key string,
	) (domain.Theme, error)

	DeleteTheme(
		ctx context.Context,
		key string,
	) error

	SaveThemes(
		ctx context.Context,
		key string,
		themes []domain.Theme,
	) error

	GetThemes(
		ctx context.Context,
		key string,
	) ([]domain.Theme, error)

	PatchTheme(
		ctx context.Context,
		id int,
		patch domain.ThemePatch,
	) error
}

type ThemeRepositoryPostgres interface {
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
	themeRepositoryPostgres ThemeRepositoryPostgres,
	themeRepositoryRedis ThemeRepositoryRedis,
) *ThemeService {
	return &ThemeService{
		themeRepositoryPostgres: themeRepositoryPostgres,
		themeRepositoryRedis:    themeRepositoryRedis,
	}
}
