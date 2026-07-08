package themes_transport_http

import (
	"fmt"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
	core_http_types "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/types"
)

type PatchThemeRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Subject     core_http_types.Nullable[string] `json:"subject"`
}

func (r *PatchThemeRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("'Title' can't be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("'Title' must be between 1 and 100 symbols")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("'Description' must be between 1 and 1000 symbols")
			}
		}
	}

	if r.Subject.Set {
		if r.Subject.Value == nil {
			return fmt.Errorf("Subject' can't be NULL")
		}

		subjectLen := len([]rune(*r.Subject.Value))
		if subjectLen < 1 || subjectLen > 1000 {
			return fmt.Errorf("'Subject' must be between 1 and 1000 symbols")
		}
	}

	return nil
}

type PatchThemeResponse ThemeDTOResponse

func (h *ThemeHTTPHandler) PatchTheme(rw http.ResponseWriter, r *http.Request) {
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

	var request PatchThemeRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	themePatch := themePatchFromRequest(request)

	themeDomain, err := h.themeService.PatchTheme(ctx, themeID, themePatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch theme",
		)

		return
	}

	response := PatchThemeResponse(themeDTOFromDomain(themeDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func themePatchFromRequest(request PatchThemeRequest) domain.ThemePatch {
	return domain.NewThemePatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Subject.ToDomain(),
	)
}
