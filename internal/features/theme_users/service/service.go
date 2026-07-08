package themes_user_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeUserService struct {
	themeUserRepository ThemeUserRepository
}

type ThemeUserRepository interface {
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
	themeUserRepository ThemeUserRepository,
) *ThemeUserService {
	return &ThemeUserService{
		themeUserRepository: themeUserRepository,
	}
}
