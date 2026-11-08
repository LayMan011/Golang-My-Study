package themes_redis_repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestThemeRedisRoundTrip(t *testing.T) {
	ctx := context.Background()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() {
		_ = client.Close()
	})

	pool := &testRedisPool{client: client}

	description := "Complete description"
	rating := 4.75
	createdAt := time.Date(2026, 7, 13, 8, 59, 8, 379286123, time.UTC)

	want := domain.Theme{
		ID:              101,
		Version:         7,
		Title:           "Redis and Go",
		Description:     &description,
		CreatedAt:       createdAt,
		Subject:         "Databases",
		Rating:          &rating,
		AllRatings:      19,
		NumberOfRatings: 6,
		NumberOfUsers:   42,
		Price:           3500,
		Level:           "advanced",
		Duration:        "6 weeks",
		Format:          "video",
		AuthorUserID:    23,
	}

	repo := NewThemesRepository(pool)
	const key = "theme:roundtrip"

	if err := repo.SaveTheme(ctx, key, want); err != nil {
		t.Fatalf("SaveTheme() error: %v", err)
	}

	got, err := repo.GetTheme(ctx, key)
	if err != nil {
		t.Fatalf("GetTheme() error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}
