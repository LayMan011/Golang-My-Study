package users_service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func generate_users() []domain.User {
	return []domain.User{
		domain.NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "1234"),
		domain.NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "1234"),
		domain.NewUserUninitailized("user@example.com", time.Now(), []byte("1234567890"), "1234"),
	}
}

func Begin(redisErr error) (context.Context, *UserService, domain.User) {
	ctx := core_logger.ToContext(
		context.Background(),
		&core_logger.Logger{Logger: zap.NewNop()},
	)

	inputUser := domain.User{
		ID:        42,
		Version:   1,
		Email:     "user@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов Иванович",
	}

	postgres := &postgresRepoStub{users: []domain.User{}}
	redis := &redisRepoStub{err: redisErr}

	service := &UserService{
		userRepositoryPostgres: postgres,
		userRepositoryRedis:    redis,
	}

	return ctx, service, inputUser
}

func TestUserService_CreateUser(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputUser := Begin(test.cacheErr)

			created, err := service.CreateUser(ctx, inputUser)
			if err != nil {
				t.Fatalf("CreateUser() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(created, inputUser) {
				t.Errorf("CreateUser() = %v, want %v", created, inputUser)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputUser := Begin(test.cacheErr)

			created, err := service.CreateUser(ctx, inputUser)
			if err != nil {
				t.Fatalf("CreateUser() returned unexpected error: %v", err)
			}

			if err := service.DeleteUser(ctx, created.ID); err != nil {
				t.Fatalf("DeleteUser() returned unexpected error: %v", err)
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputUser := Begin(test.cacheErr)

			created, err := service.CreateUser(ctx, inputUser)
			if err != nil {
				t.Fatalf("CreateUser() returned unexpected error: %v", err)
			}

			got, err := service.GetUser(ctx, created.ID)
			if err != nil {
				t.Fatalf("GetUser() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, created) {
				t.Errorf("GetUser() = %v, want %v", got, created)
			}
		})
	}
}

func TestUserService_GetUsers(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, _ := Begin(test.cacheErr)
			users := generate_users()

			for _, user := range users {
				if _, err := service.CreateUser(ctx, user); err != nil {
					t.Fatalf("CreateUser() returned unexpected error: %v", err)
				}
			}

			got, err := service.GetUsers(ctx, nil, nil)
			if err != nil {
				t.Fatalf("GetUsers() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, users) {
				t.Errorf("GetUsers() = %v, want %v", got, users)
			}
		})
	}
}

func TestUserService_PatchUser(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputUser := Begin(test.cacheErr)

			created, err := service.CreateUser(ctx, inputUser)
			if err != nil {
				t.Fatalf("CreateUser() returned unexpected error: %v", err)
			}

			fullName := "123"
			patch := domain.UserPatch{
				Password: domain.Nullable[[]byte]{
					Set:   false,
					Value: nil,
				},
				FullName: domain.Nullable[string]{
					Set:   true,
					Value: &fullName,
				},
			}

			got, err := service.PatchUser(ctx, created.ID, patch)
			if err != nil {
				t.Fatalf("PatchUser() returned unexpected error: %v", err)
			}

			if got.FullName != fullName {
				t.Errorf("PatchUser() FullName = %q, want %q", got.FullName, fullName)
			}
		})
	}
}
