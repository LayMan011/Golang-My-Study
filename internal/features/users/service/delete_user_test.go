package users_service_test

import (
	"context"
	"errors"
	"testing"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	users_service "github.com/LayMan011/Golang-My-Study/internal/features/users/service"
)

type conflictPostgresRepository struct {
	users_service.UserRepositoryPostgres
	deleteErr error
}

func (r *conflictPostgresRepository) DeleteUser(
	ctx context.Context,
	id int,
) error {
	return r.deleteErr
}

type sessionsRedisRepository struct {
	users_service.UserRepositoryRedis
	sessions                map[string]bool
	deleteSessionsCallCount int
}

func (r *sessionsRedisRepository) DeleteUserSessions(
	ctx context.Context,
	userID int,
) error {
	r.deleteSessionsCallCount++
	clear(r.sessions)
	return nil
}

type themeUserCachePostgresStub struct {
	users_service.ThemeUserCachePostgres
}

func (r *themeUserCachePostgresStub) GetIDsByThemeID(ctx context.Context, themeID int) ([]int, error) {
	return nil, nil
}

type themeUserCacheRedisStub struct {
	users_service.ThemeUserCacheRedis
}

func (r *themeUserCacheRedisStub) InvalidateByIDs(ctx context.Context, ids []int) error {
	return nil
}

func TestDeleteUser_PostgresConflictLeavesSessionsActive(t *testing.T) {
	postgresRepo := &conflictPostgresRepository{
		deleteErr: core_errors.ErrConflict,
	}

	redisRepo := &sessionsRedisRepository{
		sessions: map[string]bool{
			"session-1": true,
		},
	}

	service := users_service.NewUsersService(postgresRepo, redisRepo, &themeUserCachePostgresStub{}, &themeUserCacheRedisStub{})

	err := service.DeleteUser(context.Background(), 42)
	if !errors.Is(err, core_errors.ErrConflict) {
		t.Fatalf("DeleteUser() error = %v, want %v", err, core_errors.ErrConflict)
	}

	if redisRepo.deleteSessionsCallCount != 0 {
		t.Fatalf("DeleteUserSessions() called %d times, want 0", redisRepo.deleteSessionsCallCount)
	}

	if !redisRepo.sessions["session-1"] {
		t.Fatal("existing session was removed; want it to remain active")
	}
}
