package themes_user_transport_http

import (
	"fmt"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
	core_http_types "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/types"
)

type PatchThemeUserRequest struct {
	Completed core_http_types.Nullable[bool] `json:"completed"`
}

func (r *PatchThemeUserRequest) Validate() error {
	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("Completed' can't be NULL")
		}
	}

	return nil
}

type PatchThemeUserResponse ThemeUserDTOResponse

func (h *ThemeUserHTTPHandler) PatchThemeUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	themeUserID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get themeUserID path value",
		)

		return
	}

	var request PatchThemeUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	themeUserPatch := themeUserPatchFromRequest(request)

	themeDomain, err := h.themeUserService.PatchThemeUser(ctx, themeUserID, themeUserPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch theme",
		)

		return
	}

	response := PatchThemeUserResponse(themeUserDTOFromDomain(themeDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func themeUserPatchFromRequest(request PatchThemeUserRequest) domain.ThemeUserPatch {
	return domain.ThemeUserPatch{
		Completed: request.Completed.ToDomain(),
	}
}
