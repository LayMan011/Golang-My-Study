package themes_user_transport_http

import (
	"context"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_http_server "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/server"
)

type ThemeUserHTTPHandler struct {
	themeUserService ThemeUserService
}

type ThemeUserService interface {
	CreateThemeUser(
		ctx context.Context,
		themeUser domain.ThemeUser,
	) (domain.ThemeUser, error)

	GetThemesUser(
		ctx context.Context,
		userID *int,
		themeID *int,
		first *int,
		last *int,
	) ([]domain.ThemeUser, error)

	GetThemeUser(
		ctx context.Context,
		id int,
	) (domain.ThemeUser, error)

	GetThemeUserByUserID(
		ctx context.Context,
		id int,
	) (domain.ThemeUser, error)

	DeleteThemeUser(
		ctx context.Context,
		themeUserID int,
	) error

	PatchThemeUser(
		ctx context.Context,
		themeUserID int,
		themeUserPatch domain.ThemeUserPatch,
	) (domain.ThemeUser, error)
}

func NewThemeUserHTTPHandler(
	themeUserRepository ThemeUserService,
) *ThemeUserHTTPHandler {
	return &ThemeUserHTTPHandler{
		themeUserService: themeUserRepository,
	}
}

func (h *ThemeUserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/themes_user",
			Handler: h.CreateThemeUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/themes_user",
			Handler: h.GetThemesUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/themes_user/user/{id}",
			Handler: h.GetThemeUserByUserID,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/themes_user/{id}",
			Handler: h.DeleteThemeUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/themes_user/{id}",
			Handler: h.GetThemeUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/themes_user/{id}",
			Handler: h.PatchThemeUser,
		},
	}
}
