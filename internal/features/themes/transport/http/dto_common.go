package themes_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeDTOResponse struct {
	ID              int       `json:"id"`
	Version         int       `json:"version"`
	Title           string    `json:"title"`
	Description     *string   `json:"description"`
	CreateAt        time.Time `json:"create_at"`
	Subject         string    `json:"subject"`
	AllRatings      int       `json:"all_ratings"`
	NumberOfRatings int       `json:"number_of_ratings"`
	NumberOfUsers   int       `json:"number_of_users"`
	Price           int       `json:"price"`
	AuthorUserID    int       `json:"author_user_id"`
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
