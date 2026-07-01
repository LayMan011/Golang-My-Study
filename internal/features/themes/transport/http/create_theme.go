package themes_transport_http

import (
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type CreateThemeRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`

	AuthorUserID int `json:"author_user_id" validate:"required"`
}

type CreateThemeResponse ThemeDTOResponse

func (h *ThemeHTTPHandler) CreateTheme(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateThemeRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	themeDomain := domain.NewThemeUnitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	themeDomain, err := h.themeService.CreateTheme(ctx, themeDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create theme",
		)

		return
	}

	response := CreateThemeResponse(themeDTOFromDomain(themeDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
