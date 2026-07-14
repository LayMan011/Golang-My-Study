package users_transport_http

import (
	"net/http"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
}

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/register", registerHandler)
// 	mux.HandleFunc("/login", loginHandler)
// 	mux.HandleFunc("/refresh", refreshHandler)
// 	mux.HandleFunc("/protected", authMiddleware(protectedHandler))

// 	log.Println("server on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials
// 	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		return
// 	}
// 	if creds.Username == "" || creds.Password == "" {
// 		http.Error(w, "empty credentials", http.StatusBadRequest)
// 		return
// 	}
// 	if _, exists := users[creds.Username]; exists {
// 		http.Error(w, "user already exists", http.StatusConflict)
// 		return
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		http.Error(w, "could not hash password", http.StatusInternalServerError)
// 		return
// 	}

// 	users[creds.Username] = User{
// 		Username:     creds.Username,
// 		PasswordHash: string(hash),
// 	}

// 	writeJSON(w, map[string]string{"status": "registered"})
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var creds Credentials
// 	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		return
// 	}

// 	user, ok := users[creds.Username]
// 	if !ok {
// 		http.Error(w, "invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
// 		http.Error(w, "invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	pair, err := domain.GenerateTokenPair(user.Username)
// 	if err != nil {
// 		http.Error(w, "could not create tokens", http.StatusInternalServerError)
// 		return
// 	}

// 	err = h.redisService.SaveToken(ctx, request.Login, pair.AccessToken, 24*time.Hour)
// 	if err != nil {
// 		// Логируем ошибку, но не прерываем процесс (токен уже создан)
// 		log.Error("failed to save token to Redis", "error", err)
// 		// Можно вернуть ошибку клиенту, если хотите строгую политику
// 		// responseHandler.ErrorResponse(err, "failed to save session")
// 		// return
// 	}

// 	writeJSON(w, pair)
// }

// func refreshHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	refreshToken := readBearer(r)
// 	if refreshToken == "" {
// 		http.Error(w, "missing refresh token", http.StatusUnauthorized)
// 		return
// 	}

// 	claims, err := parseToken(refreshToken)
// 	if err != nil {
// 		http.Error(w, "invalid refresh token", http.StatusUnauthorized)
// 		return
// 	}

// 	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
// 		http.Error(w, "refresh token expired", http.StatusUnauthorized)
// 		return
// 	}

// 	key := "refresh:" + claims.ID
// 	userID, err := rdb.Get(ctx, key).Result()
// 	if err != nil {
// 		if errors.Is(err, redis.Nil) {
// 			http.Error(w, "refresh token revoked", http.StatusUnauthorized)
// 			return
// 		}
// 		http.Error(w, "redis error", http.StatusInternalServerError)
// 		return
// 	}

// 	if userID != claims.Subject {
// 		http.Error(w, "refresh token mismatch", http.StatusUnauthorized)
// 		return
// 	}

// 	_ = rdb.Del(ctx, key).Err()

// 	pair, err := domain.GenerateTokenPair(userID)
// 	if err != nil {
// 		http.Error(w, "could not refresh tokens", http.StatusInternalServerError)
// 		return
// 	}

// 	writeJSON(w, pair)
// }

func (h *UsersHTTPHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.userService.GetUserByLogin(ctx, request.Login)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(userDomain.Password)); err != nil {
		responseHandler.ErrorResponse(
			err,
			"incorrect password entered",
		)

		return
	}

	pair, err := domain.GenerateTokenPair(request.Login)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"could not create tokens",
		)

		return
	}

	if err := h.userService.Login(ctx, userDomain.Login, pair); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to write the token",
		)

		return
	}

	responseHandler.JSONResponse(pair, http.StatusOK)
}
