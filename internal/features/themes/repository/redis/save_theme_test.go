package themes_redis_repository

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestThemeRedisRepository(t *testing.T) {
	ctx := context.Background()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})
	t.Cleanup(func() {
		_ = client.Close()
	})

	repo := &ThemeRepository{
		pool: &testRedisPool{client: client},
	}

	const key = "theme:test-roundtrip"
	createdAt := time.Date(2026, 7, 13, 8, 59, 8, 379286000, time.UTC)

	want := domain.NewThemeUnitialized(
		"123",
		nil,
		"Русский язык",
		5000,
		"beginner",
		"3 месяца",
		"video",
		26,
	)
	want.CreatedAt = createdAt

	if err := repo.SaveTheme(ctx, key, want); err != nil {
		t.Fatalf("SaveTheme() error: %v", err)
	}

	got, err := repo.GetTheme(ctx, key)
	if err != nil {
		t.Fatalf("GetTheme() error: %v", err)
	}

	if got.Title != want.Title {
		t.Errorf("Title: got %q, want %q", got.Title, want.Title)
	}

	if got.Description != nil {
		t.Logf("got.Description value = %q", *got.Description)
	}

	if got.Subject != want.Subject {
		t.Errorf("Subject: got %q, want %q", got.Subject, want.Subject)
	}

	if got.Price != want.Price {
		t.Errorf("Price: got %d, want %d", got.Price, want.Price)
	}

	if got.Level != want.Level {
		t.Errorf("Level: got %q, want %q", got.Level, want.Level)
	}

	if got.Duration != want.Duration {
		t.Errorf("Duration: got %q, want %q", got.Duration, want.Duration)
	}

	if got.Format != want.Format {
		t.Errorf("Format: got %q, want %q", got.Format, want.Format)
	}

	if got.AuthorUserID != want.AuthorUserID {
		t.Errorf("AuthorUserID: got %d, want %d", got.AuthorUserID, want.AuthorUserID)
	}

	if got.CreatedAt != want.CreatedAt {
		t.Errorf("CreatedAt: got %s, want %s", got.CreatedAt, want.CreatedAt)
	}
}
