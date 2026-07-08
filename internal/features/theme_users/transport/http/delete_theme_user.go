package themes_user_transport_http

import (
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

func (h *ThemeUserHTTPHandler) DeleteThemeUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	themeUserID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get themeID path value",
		)

		return
	}

	if err := h.themeUserService.DeleteThemeUser(ctx, themeUserID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete themeUser",
		)

		return
	}

	responseHandler.NoContentResponse()
}
