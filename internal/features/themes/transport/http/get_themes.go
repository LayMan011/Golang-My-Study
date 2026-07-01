package themes_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetThemesResponse []ThemeDTOResponse

func (h *ThemeHTTPHandler) GetThemes(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/limit/offset query params",
		)

		return
	}

	themesDomains, err := h.themeService.GetThemes(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get themes",
		)

		return
	}

	response := GetThemesResponse(themesDTOsFromDomains(themesDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIDQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}
