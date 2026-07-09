package themes_user_transport_http

import (
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type CreateThemeUserRequest struct {
	ThemeID int `json:"theme_id" validate:"required" example:"1"`
	UserID  int `json:"user_id" validate:"required" example:"1"`
}

type CreateThemeUserResponse ThemeUserDTOResponse

// CreateThemeUser 	godoc
// @Summary 		Создать тему для пользователя
// @Description 	Создать новую тему для пользователя в системе
// @Tags 			themes_user
// @Accept 			json
// @Produce 		json
// @Param 			request body CreateThemeUserRequest true "CreateThemeUser тело запроса"
// @Success 		201 {object} CreateThemeUserResponse "Успешно созданная тема для пользователя"
// @Failure 		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 		404 {object} core_http_response.ErrorResponse "Author not found"
// @Failure 		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 			/themes_user [post]
func (h *ThemeUserHTTPHandler) CreateThemeUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateThemeUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	themeDomain := domain.NewThemeUserUnitialized(
		request.ThemeID,
		request.UserID,
	)

	themeDomain, err := h.themeUserService.CreateThemeUser(ctx, themeDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create theme",
		)

		return
	}

	response := CreateThemeUserResponse(themeUserDTOFromDomain(themeDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
