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
	Completed core_http_types.Nullable[bool] `json:"completed" example:"true"`
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

// PatchThemeUser 	godoc
// @Summary		Изменение темы пользователя
// @Description Изменение информации об уже существующем в системе теме пользователя
// @Description Ограничение: `completed` не может быть выставлен как null
// @Tags 		themes_user
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "ID изменяемой темы пользователя"
// @Param 		request body PatchThemeUserRequest true "PatchThemeUser тело запроса"
// @Success 	200 {object} PatchThemeUserResponse "Успешное изменение темы пользователя"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		404 {object} core_http_response.ErrorResponse "ThemeUser not found"
// @Failure		409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/themes_user/{id} [patch]
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
