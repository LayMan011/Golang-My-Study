package users_transport_http

import (
	"fmt"
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
	core_http_types "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Password core_http_types.Nullable[[]byte] `json:"password" swaggertype:"string" format:"base64"`
	FullName core_http_types.Nullable[string] `json:"full_name" swaggertype:"string" example:"Максим Максимович"`
}

// PatchUser 	godoc
// @Summary		Изменение пользователя
// @Description Изменение информации об уже существующем в системе пользователе
// @Description ### Логика обновления полей (Three-state logic):
// @Description 1. **Поле не передано**: `"phone_number"` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `"phone_number": "+71112223344"` - устанавливает новый номер телефона в БД
// @Description 3. **Передан null**: `"phone_number": null` - очищает поле в БД (set to NULL)
// @Description Ограничение: `full_name` не может быть выставлен как null
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "ID изменяемого пользователя"
// @Param 		request body PatchUserRequest true "PatchUser тело запроса"
// @Success 	200 {object} PatchUserResponse "Успешное изменение пользователя"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		404 {object} core_http_response.ErrorResponse "User not found"
// @Failure		409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/users/{id} [patch]
func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("'FullName' can't be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("'FullName' must be between 3 and 100 symbol")
		}
	}

	if r.Password.Set {
		if r.Password.Value == nil {
			return fmt.Errorf("'Password' can't be NULL")
		}

		passwordLen := len([]rune(string(*r.Password.Value)))
		if passwordLen < 8 || passwordLen > 70 {
			return fmt.Errorf("'Password' must be between 8 and 70 symbol")
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.userService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.Password.ToDomain(),
		request.FullName.ToDomain(),
	)
}
