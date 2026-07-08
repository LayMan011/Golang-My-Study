package themes_postgres_repository

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeModel struct {
	ID              int
	Version         int
	Title           string
	Description     *string
	CreatedAt       time.Time
	Subject         string
	Rating          *float64
	AllRatings      int
	NumberOfRatings int
	NumberOfUsers   int
	Price           int
	AuthorUserID    int
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
		themeModel.CreatedAt,
		themeModel.Subject,
		themeModel.Rating,
		themeModel.AllRatings,
		themeModel.NumberOfRatings,
		themeModel.NumberOfUsers,
		themeModel.Price,
		themeModel.AuthorUserID,
	)
}
