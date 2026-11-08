package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_middleware "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/middleware"
	users_redis_repository "github.com/LayMan011/Golang-My-Study/internal/features/users/repository/redis"
	users_service "github.com/LayMan011/Golang-My-Study/internal/features/users/service"
	users_transport_http "github.com/LayMan011/Golang-My-Study/internal/features/users/transport/http"
	"golang.org/x/crypto/bcrypt"
)

func TestDeleteUser_RevokesAllSessions(t *testing.T) {
	const password = "correct-password"

	t.Setenv("SECTER_KEY_FOR_JWT", "test-secret")
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	usersRepo, _, _ := initFeatures(ctx, TestDB, redisClient)
	usersRedisRepo := users_redis_repository.NewUsersRepository(redisClient)
	userService := users_service.NewUsersService(usersRepo, usersRedisRepo)
	usersHandler := users_transport_http.NewUsersHTTPHandler(userService)

	login := fmt.Sprintf("token-revocation-%d@example.com", time.Now().UnixNano())

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}

	user, err := usersRepo.CreateUser(ctx, domain.User{
		Email:     login,
		CreatedAt: time.Now(),
		Password:  passwordHash,
		FullName:  "Token Revocation Test",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	type session struct {
		accessToken  string
		refreshToken string
		jti          string
	}

	issueSession := func() session {
		t.Helper()

		body, err := json.Marshal(users_transport_http.LoginRequest{
			Email:    login,
			Password: password,
		})
		if err != nil {
			t.Fatalf("marshal login request: %v", err)
		}

		req := httptest.NewRequest(
			http.MethodPost,
			"/users/login",
			bytes.NewReader(body),
		)
		req.Header.Set("Content-Type", "application/json")
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		usersHandler.LoginUser(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("LoginUser status = %d, want %d; body = %q",
				rec.Code, http.StatusOK, rec.Body.String())
		}

		var tokens users_transport_http.LoginResponse
		if err := json.NewDecoder(rec.Body).Decode(&tokens); err != nil {
			t.Fatalf("decode login response: %v", err)
		}

		if tokens.AccessToken == "" {
			t.Fatal("LoginUser returned an empty access token")
		}

		if tokens.RefreshToken == "" {
			t.Fatal("LoginUser returned an empty refresh token")
		}

		claims, err := domain.ParseToken(tokens.AccessToken)
		if err != nil {
			t.Fatalf("parse access token: %v", err)
		}

		if claims.ID == "" {
			t.Fatal("access token has an empty jti")
		}

		refreshClaims, err := domain.ParseToken(tokens.RefreshToken)
		if err != nil {
			t.Fatalf("parse refresh token: %v", err)
		}

		if refreshClaims.ID == "" {
			t.Fatal("refresh token has an empty jti")
		}

		if refreshClaims.Subject != login {
			t.Errorf("refresh token subject = %q, want %q", refreshClaims.Subject, login)
		}

		return session{
			accessToken:  tokens.AccessToken,
			refreshToken: tokens.RefreshToken,
			jti:          claims.ID,
		}
	}

	sessions := []session{
		issueSession(),
		issueSession(),
	}
	if sessions[0].jti == sessions[1].jti {
		t.Fatal("two logins returned the same session ID; expected separate sessions")
	}

	for i, s := range sessions {
		pair, err := redisClient.HGetAll(ctx, "token:"+s.jti)
		if err != nil {
			t.Fatalf("read session %d from Redis: %v", i+1, err)
		}

		if pair["access_token"] != s.accessToken {
			t.Errorf("session %d access token was not saved in Redis", i+1)
		}

		if pair["refresh_token"] != s.refreshToken {
			t.Errorf("session %d refresh token was not saved in Redis", i+1)
		}
	}

	handlerCalls := 0
	protected := core_http_middleware.AuthMiddleware(redisClient)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerCalls++

			gotLogin, ok := core_http_middleware.LoginFromContext(r.Context())
			if !ok {
				t.Error("login is missing from middleware context")
			} else if gotLogin != login {
				t.Errorf("login in context = %q, want %q", gotLogin, login)
			}

			w.WriteHeader(http.StatusNoContent)
		}),
	)

	requestProtected := func(accessToken string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+accessToken)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		protected.ServeHTTP(rec, req)
		return rec
	}

	for i, s := range sessions {
		callsBefore := handlerCalls
		rec := requestProtected(s.accessToken)

		if rec.Code != http.StatusNoContent {
			t.Fatalf("session %d before deletion: status = %d, want %d; body = %q", i+1, rec.Code, http.StatusNoContent, rec.Body.String())
		}

		if handlerCalls != callsBefore+1 {
			t.Fatalf("session %d: next handler calls = %d, want 1", i+1, handlerCalls-callsBefore)
		}
	}

	if err := userService.DeleteUser(ctx, user.ID); err != nil {
		t.Fatalf("DeleteUser() error: %v", err)
	}

	for i, s := range sessions {
		callsBefore := handlerCalls
		rec := requestProtected(s.accessToken)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("session %d after deletion: status = %d, want %d; body = %q", i+1, rec.Code, http.StatusUnauthorized, rec.Body.String())
		}
		if handlerCalls != callsBefore {
			t.Errorf("next handler was called for session %d after deletion", i+1)
		}

		pair, err := redisClient.HGetAll(ctx, "token:"+s.jti)
		if err != nil {
			t.Fatalf("read session %d from Redis after deletion: %v", i+1, err)
		}

		if len(pair) != 0 {
			t.Errorf("session %d still has token data in Redis after deletion: %#v", i+1, pair)
		}
	}
}
