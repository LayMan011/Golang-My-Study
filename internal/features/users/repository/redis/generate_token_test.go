package users_redis_repository

import (
	"context"
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestGenerateToken(t *testing.T) {
	const (
		userID = "user-123"
		secret = "test-secret"
	)

	t.Setenv("SECTER_KEY_FOR_JWT", "test-secret")

	redisServer := miniredis.RunT(t)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
	t.Cleanup(func() {
		_ = redisClient.Close()
	})

	pool_sub := NewPoolStub()

	r := NewUsersRepository(pool_sub)

	first, err := r.GenerateTokenPair(context.Background(), userID)
	if err != nil {
		t.Fatalf("generate first token pair: %v", err)
	}

	second, err := r.GenerateTokenPair(context.Background(), userID)
	if err != nil {
		t.Fatalf("generate second token pair: %v", err)
	}

	firstAccess, err := domain.ParseToken(first.AccessToken)
	if err != nil {
		t.Fatalf("parse first access token: %v", err)
	}

	firstRefresh, err := domain.ParseToken(first.RefreshToken)
	if err != nil {
		t.Fatalf("parse first refresh token: %v", err)
	}

	secondAccess, err := domain.ParseToken(second.AccessToken)
	if err != nil {
		t.Fatalf("parse second access token: %v", err)
	}

	secondRefresh, err := domain.ParseToken(second.RefreshToken)
	if err != nil {
		t.Fatalf("parse second refresh token: %v", err)
	}

	type claim struct {
		name   string
		claims *domain.Claims
	}

	claims := []claim{
		claim{"first access", firstAccess},
		claim{"first refresh", firstRefresh},
		claim{"second access", secondAccess},
		claim{"second refresh", secondRefresh},
	}

	for _, item := range claims {
		if item.claims.Subject != userID {
			t.Errorf("%s sub = %q, want %q", item.name, item.claims.Subject, userID)
		}

		if item.claims.ID == "" {
			t.Errorf("%s jti is empty", item.name)
		}
	}

	seen := make(map[string]string)
	for _, item := range claims {
		if previous, exists := seen[item.claims.ID]; exists {
			t.Errorf("%s and %s have the same jti %q", previous, item.name, item.claims.ID)
			continue
		}

		seen[item.claims.ID] = item.name
	}
}
