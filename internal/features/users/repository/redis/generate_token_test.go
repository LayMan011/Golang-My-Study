package users_redis_repository

import (
	"context"
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func TestGenerateToken(t *testing.T) {
	const (
		userID = "user-123"
		secret = "test-secret"
	)

	t.Setenv("SECTER_KEY_FOR_JWT", secret)

	ctx := context.Background()
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() {
		_ = redisClient.Close()
	})

	pool := &miniredisPool{
		PoolStub: NewPoolStub(),
		client:   redisClient,
	}
	r := NewUsersRepository(pool)

	first, err := r.GenerateTokenPair(ctx, userID)
	if err != nil {
		t.Fatalf("generate first token pair: %v", err)
	}

	second, err := r.GenerateTokenPair(ctx, userID)
	if err != nil {
		t.Fatalf("generate second token pair: %v", err)
	}

	parseRefreshClaims := func(tokenString string) *domain.Claims {
		t.Helper()

		claims := &domain.Claims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(_ *jwt.Token) (any, error) {
				return []byte(secret), nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)
		if err != nil {
			t.Fatalf("parse refresh token: %v", err)
		}

		if !token.Valid {
			t.Fatal("refresh token is invalid")
		}

		return claims
	}

	firstRefresh := parseRefreshClaims(first.RefreshToken)
	secondRefresh := parseRefreshClaims(second.RefreshToken)

	if firstRefresh.ID == secondRefresh.ID {
		t.Fatalf("expected distinct refresh jti values, got %q", firstRefresh.ID)
	}

	keys, err := redisClient.Keys(ctx, "refresh:*").Result()
	if err != nil {
		t.Fatalf("list refresh-session keys: %v", err)
	}

	if len(keys) != 2 {
		t.Fatalf("expected 2 refresh-session keys in Redis, got %d: %v", len(keys), keys)
	}

	for _, claims := range []*domain.Claims{firstRefresh, secondRefresh} {
		key := "refresh:" + claims.ID

		got, err := redisClient.Get(ctx, key).Result()
		if err != nil {
			t.Fatalf("get refresh session %q from Redis: %v", key, err)
		}

		if got != userID {
			t.Errorf("refresh session %q contains user ID %q, want %q", key, got, userID)
		}
	}
}

func TestGenerateToken_SavesRefreshSessions(t *testing.T) {
	const (
		userID = "user-123"
		secret = "test-secret"
	)

	t.Setenv("SECTER_KEY_FOR_JWT", secret)

	pool := &recordingPool{values: make(map[string]string)}
	repo := NewUsersRepository(pool)
	ctx := context.Background()

	firstPair, err := repo.GenerateTokenPair(ctx, userID)
	if err != nil {
		t.Fatalf("GenerateTokenPair() first call returned error: %v", err)
	}

	secondPair, err := repo.GenerateTokenPair(ctx, userID)
	if err != nil {
		t.Fatalf("GenerateTokenPair() second call returned error: %v", err)
	}

	refreshJTI := func(tokenString string) string {
		t.Helper()

		claims := &domain.Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (any, error) {
			return []byte(secret), nil
		})
		if err != nil {
			t.Fatalf("failed to parse refresh token: %v", err)
		}

		return claims.ID
	}

	firstKey := "refresh:" + refreshJTI(firstPair.RefreshToken)
	secondKey := "refresh:" + refreshJTI(secondPair.RefreshToken)

	if firstKey == secondKey {
		t.Fatalf("expected distinct refresh-session keys, got %q for both tokens", firstKey)
	}

	if len(pool.values) != 2 {
		t.Fatalf("expected 2 saved refresh sessions, got %d", len(pool.values))
	}

	for _, key := range []string{firstKey, secondKey} {
		if got, ok := pool.values[key]; !ok {
			t.Errorf("expected refresh session %q to be saved", key)
		} else if got != userID {
			t.Errorf("refresh session %q: expected user ID %q, got %q", key, userID, got)
		}
	}
}
