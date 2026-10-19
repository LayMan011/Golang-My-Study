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
	CreatedAt    time.Time
	Subject      string
	Price        int
	Completed    bool
	CompletedAt  *time.Time
	Level        string
	Duration     string
	Format       string
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
		*themeModel.CompletedAt,
		themeModel.Subject,
		new(float64),
		0,
		0,
		0,
		themeModel.Price,
		themeModel.Level,
		themeModel.Duration,
		themeModel.Format,
		themeModel.AuthorUserID,
	)
}
