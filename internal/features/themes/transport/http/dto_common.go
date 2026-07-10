package themes_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeDTOResponse struct {
	ID              int       `json:"id" example:"13"`
	Version         int       `json:"version" example:"3"`
	Title           string    `json:"title" example:"Математика"`
	Description     *string   `json:"description" example:"Подготовка к ЕГЭ по профильной математике"`
	CreateAt        time.Time `json:"create_at" example:"Профильная математика"`
	Subject         string    `json:"subject" example:"2026-02-26T10:30:00Z"`
	AllRatings      int       `json:"all_ratings" example:"0"`
	NumberOfRatings int       `json:"number_of_ratings" example:"0"`
	NumberOfUsers   int       `json:"number_of_users" example:"0"`
	Price           int       `json:"price" example:"50"`
	AuthorUserID    int       `json:"author_user_id" example:"2"`
}

func themeDTOFromDomain(theme domain.Theme) ThemeDTOResponse {
	return ThemeDTOResponse{
		ID:              theme.ID,
		Version:         theme.Version,
		Title:           theme.Title,
		Description:     theme.Description,
		CreateAt:        theme.CreatedAt,
		Subject:         theme.Subject,
		AllRatings:      theme.AllRatings,
		NumberOfRatings: theme.NumberOfRatings,
		NumberOfUsers:   theme.NumberOfUsers,
		Price:           theme.Price,
		AuthorUserID:    theme.AuthorUserID,
	}
}

func themesDTOsFromDomains(themes []domain.Theme) []ThemeDTOResponse {
	dtos := make([]ThemeDTOResponse, len(themes))
	for i, theme := range themes {
		dtos[i] = themeDTOFromDomain(theme)
	}

	return dtos
}
