package themes_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreateAt     time.Time  `json:"create_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	Percentages  int        `json:"percentages"`
	AuthorUserID int        `json:"author_user_id"`
}

func themeDTOFromDomain(theme domain.Theme) ThemeDTOResponse {
	return ThemeDTOResponse{
		ID:           theme.ID,
		Version:      theme.Version,
		Title:        theme.Title,
		Description:  theme.Description,
		Completed:    theme.Completed,
		CreateAt:     theme.CreatedAt,
		CompletedAt:  theme.CompletedAt,
		Percentages:  theme.Percentages,
		AuthorUserID: theme.AuthorUserID,
	}
}

func themesDTOsFromDomains(themes []domain.Theme) []ThemeDTOResponse {
	dtos := make([]ThemeDTOResponse, len(themes))
	for i, theme := range themes {
		dtos[i] = themeDTOFromDomain(theme)
	}

	return dtos
}
