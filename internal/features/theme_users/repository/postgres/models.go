package themes_user_postgres_repository

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeUserModel struct {
	ID          int
	Version     int
	Completed   bool
	AdditionAt  time.Time
	CompletedAt *time.Time
	Percentages int
	ThemeID     int
	UserID      int
}

func themesUserDomainFromModels(themeModels []ThemeUserModel) []domain.ThemeUser {
	domains := make([]domain.ThemeUser, len(themeModels))
	for i, model := range themeModels {
		domains[i] = themeUserDomainFromModel(model)
	}

	return domains
}

func themeUserDomainFromModel(themeModel ThemeUserModel) domain.ThemeUser {
	return domain.NewThemeUser(
		themeModel.ID,
		themeModel.Version,
		themeModel.Completed,
		themeModel.AdditionAt,
		themeModel.CompletedAt,
		themeModel.Percentages,
		themeModel.ThemeID,
		themeModel.UserID,
	)
}
