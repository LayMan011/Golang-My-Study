package themes_user_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeUserService struct {
	hash_once                   map[string]bool
	hash_many                   map[string]bool
	themeUserRepositoryPostgres ThemeUserRepositoryPostgres
	themeUserRepositoryRedis    ThemeUserRepositoryRedis
}

type ThemeUserRepositoryRedis interface {
	SaveThemeUser(
		ctx context.Context,
		key string,
		value domain.ThemeUser,
	) error

	GetThemeUser(
		ctx context.Context,
		key string,
	) (domain.ThemeUser, error)

	DeleteThemeUser(
		ctx context.Context,
		key string,
	) error

	PatchThemeUser(
		ctx context.Context,
		id int,
		patch domain.ThemeUserPatch,
	) error

	SaveThemesUser(
		ctx context.Context,
		key string,
		themesUser []domain.ThemeUser,
	) error

	GetThemesUser(
		ctx context.Context,
		key string,
	) ([]domain.ThemeUser, error)
}

type ThemeUserRepositoryPostgres interface {
	CreateThemeUser(
		ctx context.Context,
		themeUser domain.ThemeUser,
	) (domain.ThemeUser, error)

	GetThemesUser(
		ctx context.Context,
		userID *int,
		themeID *int,
		first *int,
		last *int,
	) ([]domain.ThemeUser, error)

	GetThemeUser(
		ctx context.Context,
		id int,
	) (domain.ThemeUser, error)

	GetThemeUserByUserID(
		ctx context.Context,
		id int,
	) (domain.ThemeUser, error)

	DeleteThemeUser(
		ctx context.Context,
		themeUserID int,
	) error

	PatchThemeUser(
		ctx context.Context,
		id int,
		themeUser domain.ThemeUser,
	) (domain.ThemeUser, error)
}

func NewThemeUserService(
	themeUserRepositoryPostgres ThemeUserRepositoryPostgres,
	themeUserRepositoryRedis ThemeUserRepositoryRedis,
) *ThemeUserService {
	return &ThemeUserService{
		hash_once:                   make(map[string]bool),
		hash_many:                   make(map[string]bool),
		themeUserRepositoryPostgres: themeUserRepositoryPostgres,
		themeUserRepositoryRedis:    themeUserRepositoryRedis,
	}
}
