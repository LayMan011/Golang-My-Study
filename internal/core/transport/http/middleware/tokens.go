package core_http_middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(pool core_redis_pool.Pool, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		login, err := parseLoginFromToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), pool.OpTimeout())
		defer cancel()

		key := fmt.Sprintf("jwt:%s", login)
		storedToken, err := pool.GetResult(ctx, key)
		if err != nil {
			if err == redis.Nil {
				http.Error(w, "session expired or not found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if *storedToken != tokenString {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctxWithLogin := context.WithValue(r.Context(), "login", login)
		r = r.WithContext(ctxWithLogin)

		next.ServeHTTP(w, r)
	}
}

func parseLoginFromToken(tokenString string) (string, error) {
	claims, err := domain.ParseToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return "", fmt.Errorf("token expired")
	}

	return claims.Login, nil
}
