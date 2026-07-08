package themes_user_transport_http

import (
	"fmt"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetThemesUserResponse []domain.ThemeUser

func (h *ThemeUserHTTPHandler) GetThemesUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, themeID, first, last, err := getUserIDThemeIDFirstLastQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userId/first/last query params",
		)

		return
	}

	themesUserDomain, err := h.themeUserService.GetThemesUser(ctx, userID, themeID, first, last)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed ro get themes_user",
		)

		return
	}

	response := themesUserDTOsFromDomains(themesUserDomain)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDThemeIDFirstLastQueryParams(r *http.Request) (*int, *int, *int, *int, error) {
	const (
		userIDQueryParamKey  = "user_id"
		themeIDQueryParamKey = "theme_id"
		firstQueryParamKey   = "first"
		lastQueryParamKey    = "last"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	themeID, err := core_http_request.GetIntQueryParam(r, themeIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	first, err := core_http_request.GetIntQueryParam(r, firstQueryParamKey)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("get 'first' query param: %w", err)
	}

	last, err := core_http_request.GetIntQueryParam(r, lastQueryParamKey)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("get 'last' query param: %w", err)
	}

	return userID, themeID, first, last, nil
}
