package users_transport_http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=4,max=100"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}

type LoginResponse struct {
	UserId       string `json:"user_id" example:"4"`
	AccessToken  string `json:"access_token" example:""`
	RefreshToken string `json:"refresh_token" example:""`
}

type contextKey string

const (
	ctxKeyTokenPair contextKey = "tokenPair"
)

func (h *UsersHTTPHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		fmt.Printf("DEBUG: Получено из JSON -> Email: '%s', Password: '%s'\n", request.Email, request.Password)
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.userService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(request.Password)); err != nil {
		responseHandler.ErrorResponse(
			err,
			"incorrect password entered",
		)

		return
	}

	pair, err := h.userService.GenerateTokenPair(ctx, request.Email)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"could not create tokens",
		)

		return
	}

	if err := h.userService.Login(ctx, userDomain.ID, pair); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to write the token",
		)

		return
	}

	response := LoginResponse{
		UserId:       strconv.Itoa(userDomain.ID),
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
	}

	ctx = context.WithValue(ctx, ctxKeyTokenPair, pair)
	r = r.WithContext(ctx)

	responseHandler.JSONResponse(response, http.StatusOK)
}
