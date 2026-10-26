package users_service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
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

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{users: []domain.User{}}
	redis := &redisRepoStub{err: redisErr}

	service := &UserService{
		userRepositoryPostgres: postgres,
		userRepositoryRedis:    redis,
	}

	return ctx, service, inputUser
}

func TestUserService_CreateUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	empty_user := domain.User{}

	created, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(created, empty_user) {
		t.Errorf("CreateUser() = %v, want %v", created, empty_user)
	}

	if !reflect.DeepEqual(created, inputUser) {
		t.Errorf("CreateUser() = %v, want %v", created, inputUser)
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	created, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	err = service.DeleteUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteUser() returned unexpected error: %v", err)
	}
}

func TestUserService_GetUser(t *testing.T) {
	ctx, service, inputUser := Begin()

	empty_user := domain.User{}

	created, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	get_user, err := service.GetUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetUser() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_user, empty_user) {
		t.Errorf("GetUser() = %v, want %v", get_user, empty_user)
	}

	if !reflect.DeepEqual(get_user, created) {
		t.Errorf("GetUser() = %v, want %v", get_user, created)
	}
}

func TestUserService_GetUsers(t *testing.T) {
	ctx, service, _ := Begin()

	empty_users := []domain.User{}

	users := generate_users()

	for _, user := range users {
		_, err := service.CreateUser(ctx, user)
		if err != nil {
			t.Fatalf("CreateUser() returned unexpected error: %v", err)
		}
	}

	get_users, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_users, empty_users) {
		t.Errorf("GetUsers() = %v, want %v", get_users, empty_users)
	}

	if !reflect.DeepEqual(get_users, users) {
		t.Errorf("GetUsers() = %v, want %v", get_users, users)
	}
}

func TestUserService_PatchUser(t *testing.T) {
	ctx, service, user := Begin()

	created, err := service.CreateUser(ctx, user)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

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

	user_patch, err := service.PatchUser(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchUser() returned unexpected error: %v", err)
	}

	if user_patch.FullName != fullname {
		t.Errorf("PatchTheme() FullName = %q, want %q", user_patch.FullName, fullname)
	}
}
