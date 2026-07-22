package core_http_middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

type contextKey string

const ctxKeyLogin contextKey = "login"

func AuthMiddleware(pool core_redis_pool.Pool) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}
			accessToken := parts[1]

			// 2. Парсим логин из токена
			login, err := parseLoginFromToken(accessToken)
			if err != nil || login == "" {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()

			// 3. Читаем TokenPair из Redis через GetToken
			key := fmt.Sprintf("jwt:%s", login)

			tokenPair, err := getToken(ctx, key, pool)
			if err != nil {
				if errors.Is(err, core_errors.ErrNotFound) {
					http.Error(w, "session expired or not found", http.StatusUnauthorized)
					return
				}
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			// 4. Проверяем, что accessToken совпадает с тем, что в Redis
			if tokenPair.AccessToken != accessToken {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// 5. Кладём логин (и при желании другие данные) в контекст
			ctxWithLogin := context.WithValue(ctx, ctxKeyLogin, login)
			next.ServeHTTP(w, r.WithContext(ctxWithLogin))
		})
	}
}

func getToken(ctx context.Context, key string, pool core_redis_pool.Pool) (domain.TokenPair, error) {
	data, err := pool.HGetAll(ctx, key)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("get token from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.TokenPair{}, core_errors.ErrNotFound
	}

	var t domain.TokenPair
	v := reflect.ValueOf(&t).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("redis")
		if tag == "" || tag == "-" {
			continue
		}

		raw, ok := data[tag]
		if !ok {
			continue
		}

		fv := v.Field(i)
		if !fv.CanSet() {
			continue
		}

		fv.SetString(raw)
	}

	return t, nil
}

func parseLoginFromToken(ID string) (string, error) {
	claims, err := domain.ParseToken(ID)
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return "", fmt.Errorf("token expired")
	}

	return claims.ID, nil
}
