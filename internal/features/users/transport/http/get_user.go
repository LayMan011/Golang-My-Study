package users_transport_http

import (
	"net/http"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHander := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHander.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHander.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHander.JSONResponse(response, http.StatusOK)
}
