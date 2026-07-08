package themes_user_transport_http

import (
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type ThemeUserDTOResponse struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Completed   bool       `json:"completed"`
	AdditionAt  time.Time  `json:"addition_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Percentages int        `json:"percentages"`
	ThemeID     int        `json:"theme_id"`
	UserID      int        `json:"user_id"`
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
