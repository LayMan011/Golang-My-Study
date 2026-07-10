package themes_user_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeUserDTOResponse struct {
	ID          int        `json:"id" example:"1"`
	Version     int        `json:"version" example:"2"`
	Completed   bool       `json:"completed" example:"false"`
	AdditionAt  time.Time  `json:"addition_at" example:"2026-02-26T10:30:00Z"`
	CompletedAt *time.Time `json:"completed_at" example:"null"`
	Percentages int        `json:"percentages" example:"0"`
	ThemeID     int        `json:"theme_id" example:"3"`
	UserID      int        `json:"user_id" example:"4"`
}

func themeUserDTOFromDomain(theme domain.ThemeUser) ThemeUserDTOResponse {
	return ThemeUserDTOResponse{
		ID:          theme.ID,
		Version:     theme.Version,
		Completed:   theme.Completed,
		AdditionAt:  theme.AdditionAt,
		CompletedAt: theme.CompletedAt,
		Percentages: theme.Percentages,
		ThemeID:     theme.ThemeID,
		UserID:      theme.UserID,
	}
}

func themesUserDTOsFromDomains(themes []domain.ThemeUser) []ThemeUserDTOResponse {
	dtos := make([]ThemeUserDTOResponse, len(themes))
	for i, theme := range themes {
		dtos[i] = themeUserDTOFromDomain(theme)
	}

	return dtos
}
