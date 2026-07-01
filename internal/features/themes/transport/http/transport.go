package themes_transport_http

import (
	"context"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_http_server "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/server"
)

type ThemeHTTPHandler struct {
	themeService ThemeService
}

type ThemeService interface {
	CreateTheme(
		ctx context.Context,
		theme domain.Theme,
	) (domain.Theme, error)

	GetThemes(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Theme, error)

	GetTheme(
		ctx context.Context,
		themeID int,
	) (domain.Theme, error)

	DeleteTheme(
		ctx context.Context,
		id int,
	) error

	PatchTheme(
		ctx context.Context,
		id int,
		patch domain.ThemePatch,
	) (domain.Theme, error)
}

func NewThemesHTTPHandler(
	themeService ThemeService,
) *ThemeHTTPHandler {
	return &ThemeHTTPHandler{
		themeService: themeService,
	}
}

func (h *ThemeHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/themes",
			Handler: h.CreateTheme,
		},
		{
			Method:  http.MethodGet,
			Path:    "/themes",
			Handler: h.GetThemes,
		},
		{
			Method:  http.MethodGet,
			Path:    "/themes/{id}",
			Handler: h.GetTheme,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/themes/{id}",
			Handler: h.DeleteTheme,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/themes/{id}",
			Handler: h.PatchTheme,
		},
	}
}
