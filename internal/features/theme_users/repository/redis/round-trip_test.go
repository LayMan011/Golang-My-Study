package themes_user_redis_repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestThemeUserRedisRoundTrip(t *testing.T) {
	t.Run("completed_at nil", func(t *testing.T) {
		ctx := context.Background()
		pool := newRoundTripRedisPool(t)

		want := domain.ThemeUser{
			ID:               303,
			Version:          4,
			Completed:        false,
			AdditionAt:       time.Date(2026, 1, 2, 3, 4, 5, 123456789, time.UTC),
			CompletedAt:      nil,
			Percentages:      35,
			TotalLessons:     20,
			CompletedLessons: 7,
			ThemeID:          101,
			UserID:           202,
		}

		assertThemeUserRoundTrip(t, ctx, pool, "theme-user:roundtrip-nil", want)
	})

	t.Run("completed_at is set", func(t *testing.T) {
		ctx := context.Background()
		pool := newRoundTripRedisPool(t)

		completedAt := time.Date(2026, 3, 4, 5, 6, 7, 135792468, time.UTC)
		want := domain.ThemeUser{
			ID:               304,
			Version:          5,
			Completed:        true,
			AdditionAt:       time.Date(2026, 2, 3, 4, 5, 6, 246801357, time.UTC),
			CompletedAt:      &completedAt,
			Percentages:      100,
			TotalLessons:     20,
			CompletedLessons: 20,
			ThemeID:          102,
			UserID:           203,
		}

		assertThemeUserRoundTrip(t, ctx, pool, "theme-user:roundtrip-completed", want)
	})
}

func assertThemeUserRoundTrip(
	t *testing.T,
	ctx context.Context,
	pool *testRedisPool,
	key string,
	want domain.ThemeUser,
) {
	t.Helper()

	repo := NewThemesUserRepository(pool)

	if err := repo.SaveThemeUser(ctx, key, want); err != nil {
		t.Fatalf("SaveThemeUser() error: %v", err)
	}

	got, err := repo.GetThemeUser(ctx, key)
	if err != nil {
		t.Fatalf("GetThemeUser() error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}

func newRoundTripRedisPool(t *testing.T) *testRedisPool {
	t.Helper()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() {
		_ = client.Close()
	})

	return &testRedisPool{client: client}
}
