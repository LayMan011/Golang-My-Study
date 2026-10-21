package users_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func Begin() (context.Context, *UserService, domain.User) {
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
	createdUser := inputUser

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{user: createdUser}
	redis := &redisRepoStub{err: redisErr}

	service := &UserService{
		userRepositoryPostgres: postgres,
		userRepositoryRedis:    redis,
	}

	return ctx, service, inputUser
}

func TestUserService_CreateUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	_, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	err := service.DeleteUser(ctx, inputUser.ID)
	if err != nil {
		t.Fatalf("DeleteUser() returned unexpected error: %v", err)
	}
}

func TestUserService_GetUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	_, err := service.GetUser(ctx, inputUser.ID)
	if err != nil {
		t.Fatalf("GetUser() returned unexpected error: %v", err)
	}
}

func TestUserService_GetUsers(t *testing.T) {
	ctx, service, _ := Begin()

	_, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}
}

func TestUserService_PatchUser(t *testing.T) {
	ctx, service, user := Begin()

	password := []byte("123456789")
	fullname := "123"
	patch := domain.UserPatch{
		Password: domain.Nullable[[]byte]{
			Set:   false,
			Value: &password,
		},
		FullName: domain.Nullable[string]{
			Set:   true,
			Value: &fullname,
		},
	}

	_, err := service.PatchUser(ctx, user.ID, patch)
	if err != nil {
		t.Fatalf("PatchUsers() returned unexpected error: %v", err)
	}
}
