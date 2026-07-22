package themes_user_transport_http

import (
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetThemeUserByUserIDResponse ThemeUserDTOResponse

// GetThemeUser 	godoc
// @Summary			Получение темы пользователя
// @Description 	Получение конкретной темы пользователя в системе по ее ID
// @Tags			themes_user
// @Produce 		json
// @Param 			id path int true "ID получаемой темы пользователя"
// @Success 		200 {object} GetThemeUserResponse "Тема пользователя успешно найдена"
// @Failure 		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure			404 {object} core_http_response.ErrorResponse "ThemeUser not found"
// @Failure 		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 			/themes_user/{id} [get]
func (h *ThemeUserHTTPHandler) GetThemeUserByUserID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get themeUserID path value",
		)

		return
	}

	themeUserDomain, err := h.themeUserService.GetThemeUserByUserID(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get theme_user",
		)

		return
	}

	response := GetThemeUserByUserIDResponse(themeUserDTOFromDomain(themeUserDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
