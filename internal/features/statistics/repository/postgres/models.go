package statistics_postgres_repository

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	Percentages  int
	AuthorUserID int
}

func themesDomainFromModels(themeModels []ThemeModel) []domain.Theme {
	domains := make([]domain.Theme, len(themeModels))
	for i, model := range themeModels {
		domains[i] = themeDomainFromModel(model)
	}

	return domains
}

func themeDomainFromModel(themeModel ThemeModel) domain.Theme {
	return domain.NewTheme(
		themeModel.ID,
		themeModel.Version,
		themeModel.Title,
		themeModel.Description,
		themeModel.Completed,
		themeModel.CreatedAt,
		themeModel.CompletedAt,
		themeModel.Percentages,
		themeModel.AuthorUserID,
	)
}
