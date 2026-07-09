package themes_transport_http

import (
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

// DeleteTheme 	godoc
// @Summary		Удаление темы
// @Description Удаление существующей темы в системе по ее ID
// @Tags 		themes
// @Param 		id path int true "ID удаляемой темы"
// @Success 	204 "Успешное удаление темы"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		404 {object} core_http_response.ErrorResponse "Theme not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/themes/{id} [delete]
func (h *ThemeHTTPHandler) DeleteTheme(rw http.ResponseWriter, r *http.Request) {
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

	if err := h.themeService.DeleteTheme(ctx, themeID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete theme",
		)

		return
	}

	responseHandler.NoContentResponse()
}
