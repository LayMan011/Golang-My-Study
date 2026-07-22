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
	Title       core_http_types.Nullable[string] `json:"title" swaggertype:"string" example:"Биология"`
	Description core_http_types.Nullable[string] `json:"description" swaggertype:"string" example:"Подготовка к ЕГЭ по Биологии"`
	Subject     core_http_types.Nullable[string] `json:"subject" swaggertype:"string" example:"Биология"`
	Level       core_http_types.Nullable[string] `json:"level" swaggertype:"string" example:"beginner"`
	Duration    core_http_types.Nullable[string] `json:"duration" swaggertype:"string" example:"3 месяца"`
	Format      core_http_types.Nullable[string] `json:"format" swaggertype:"string" example:"video"`
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

	if r.Level.Set {
		if r.Level.Value == nil {
			return fmt.Errorf("'Level' can't be NULL")
		}

		if *r.Level.Value != "beginner" && *r.Level.Value != "advanced" && *r.Level.Value != "intermediate" {
			return fmt.Errorf("'Level' must be 'beginner' or 'advanced' or 'intermediate'")
		}
	}

	if r.Duration.Set {
		if r.Duration.Value != nil {
			durationLen := len([]rune(*r.Duration.Value))
			if durationLen < 1 || durationLen > 40 {
				return fmt.Errorf("'Duration' must be between 1 and 40 symbols")
			}
		}
	}

	if r.Format.Set {
		if r.Format.Value == nil {
			return fmt.Errorf("'Format' can't be NULL")
		}

		if *r.Format.Value != "video" && *r.Format.Value != "text" && *r.Format.Value != "mixed" {
			return fmt.Errorf("'Format' must be 'video' or 'text' or 'mixed'")
		}
	}

	return nil
}

type PatchThemeResponse ThemeDTOResponse

// PatchTheme 	godoc
// @Summary		Изменение темы
// @Description Изменение информации об уже существующем в системе теме
// @Description ### Логика обновления полей (Three-state logic):
// @Description 1. **Поле не передано**: `"description"` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `"description": "Пойти на  тренировку"` - устанавливает новое описание в БД
// @Description 3. **Передан null**: `"description": null` - очищает поле в БД (set to NULL)
// @Description Ограничение: `title` и `subject` не может быть выставлен как null
// @Tags 		themes
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "ID изменяемой темы"
// @Param 		request body PatchThemeRequest true "PatchTheme тело запроса"
// @Success 	200 {object} PatchThemeResponse "Успешное изменение темы"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		404 {object} core_http_response.ErrorResponse "Theme not found"
// @Failure		409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/themes/{id} [patch]
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
		request.Level.ToDomain(),
		request.Duration.ToDomain(),
		request.Format.ToDomain(),
	)
}
