package themes_user_redis_repository

import (
	"context"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestSaveThemeUser_RemovesCompletedAtWhenNil(t *testing.T) {
	ctx := context.Background()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() {
		_ = client.Close()
	})

	repo := &ThemeUserRepository{pool: &testRedisPool{client: client}}
	const key = "theme-user:save-nil-completed-at"

	completedAt := time.Date(2026, 3, 4, 5, 6, 7, 0, time.UTC)
	value := domain.ThemeUser{
		ID:               1,
		Version:          1,
		Completed:        true,
		AdditionAt:       time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC),
		CompletedAt:      &completedAt,
		Percentages:      100,
		TotalLessons:     10,
		CompletedLessons: 10,
		ThemeID:          2,
		UserID:           3,
	}

	if err := repo.SaveThemeUser(ctx, key, value); err != nil {
		t.Fatalf("first SaveThemeUser() error: %v", err)
	}

	if !client.HExists(ctx, key, "completed_at").Val() {
		t.Fatal("completed_at should exist after saving a non-nil value")
	}

	value.CompletedAt = nil

	if err := repo.SaveThemeUser(ctx, key, value); err != nil {
		t.Fatalf("second SaveThemeUser() error: %v", err)
	}

	if client.HExists(ctx, key, "completed_at").Val() {
		t.Error("completed_at should have been deleted when saving nil")
	}
}
