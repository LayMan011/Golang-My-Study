package themes_transport_http

import (
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetThemeResponse ThemeDTOResponse

// GetTheme 	godoc
// @Summary		Получение темы
// @Description Получение конкретной темы в системе по ее ID
// @Tags		themes
// @Produce 	json
// @Param 		id path int true "ID получаемой темы"
// @Success 	200 {object} GetThemeResponse "Тема успешно найдена"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		404 {object} core_http_response.ErrorResponse "Theme not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/themes/{id} [get]
func (h *ThemeHTTPHandler) GetTheme(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	themeID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get themeID path value",
		)

		return
	}

	themeDomain, err := h.themeService.GetTheme(ctx, themeID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get theme",
		)

		return
	}

	response := GetThemeResponse(themeDTOFromDomain(themeDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
