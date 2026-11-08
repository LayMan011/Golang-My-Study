package users_redis_repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestUserRedisRoundTrip(t *testing.T) {
	ctx := context.Background()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() {
		_ = client.Close()
	})

	pool := &testRedisPool{client: client}

	createdAt := time.Date(2025, 11, 8, 12, 34, 56, 987654321, time.UTC)
	want := domain.User{
		ID:        202,
		Version:   9,
		Email:     "user@example.test",
		CreatedAt: createdAt,
		Password:  []byte("hashed-password-bytes"),
		FullName:  "Taylor Example",
	}

	repo := NewUsersRepository(pool)
	const key = "user:roundtrip"

	if err := repo.SaveUser(ctx, key, want); err != nil {
		t.Fatalf("SaveUser() error: %v", err)
	}

	got, err := repo.GetUser(ctx, key)
	if err != nil {
		t.Fatalf("GetUser() error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}
