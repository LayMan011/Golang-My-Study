package users_service

import (
	"context"
	"errors"
	"testing"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
)

type conflictPostgresRepository struct {
	UserRepositoryPostgres
	deleteErr error
}

func (r *conflictPostgresRepository) DeleteUser(
	ctx context.Context,
	id int,
) error {
	return r.deleteErr
}

type sessionsRedisRepository struct {
	UserRepositoryRedis
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

func TestDeleteUser_PostgresConflictLeavesSessionsActive(t *testing.T) {
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	postgresRepo := &conflictPostgresRepository{
		deleteErr: core_errors.ErrConflict,
	}

	redisRepo := &sessionsRedisRepository{
		sessions: map[string]bool{
			"session-1": true,
		},
	}

	service := NewUsersService(postgresRepo, redisRepo, &themeUserCachePostgresStub{}, &themeUserCacheRedisStub{})

	err = service.DeleteUser(ctx, 42)
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

func TestUserService_DeleteUser_GetIDsErrorDoesNotDeleteUser(t *testing.T) {
	const userID = 123

	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	wantErr := errors.New("failed to get theme_user IDs")
	themeUserPostgres := &themeUserCachePostgresStub{
		err: wantErr,
	}

	service := &UserService{
		themeUserCachePostgres: themeUserPostgres,
	}

	err = service.DeleteUser(ctx, userID)

	if !errors.Is(err, wantErr) {
		t.Fatalf("DeleteUser() error = %v, want wrapped error %v", err, wantErr)
	}

	if themeUserPostgres.getIDsByThemeIDCalls != 1 {
		t.Errorf("GetIDsByUserID calls = %d, want 1", themeUserPostgres.getIDsByThemeIDCalls)
	}
}
