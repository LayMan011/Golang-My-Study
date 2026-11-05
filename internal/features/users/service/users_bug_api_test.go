package users_service

import (
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin(nil)

	users_generate := generate_users()
	for _, user := range users_generate {
		if _, err := service.CreateUser(ctx, user); err != nil {
			t.Fatalf("CreateUser() returned unexpected error: %v", err)
		}
	}

	offset := 1
	users_query, err := service.GetUsers(ctx, nil, &offset)
	if err != nil {
		t.Fatalf("GetUsers() with query params returned unexpected error: %v", err)
	}

	users, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	if len(users_query) == len(users) {
		t.Fatalf("GetUsers with query params writes to the cache, just like for all users")
	}
}

func TestBug14_B(t *testing.T) {
	ctx, service, _ := Begin(nil)

	users_generate := generate_users()
	for _, user := range users_generate {
		if _, err := service.CreateUser(ctx, user); err != nil {
			t.Fatalf("CreateUser() returned unexpected error: %v", err)
		}
	}

	limit := 0
	users_query, err := service.GetUsers(ctx, &limit, nil)
	if err != nil {
		t.Fatalf("GetUsers() with query params returned unexpected error: %v", err)
	}

	users, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	if len(users_query) == len(users) {
		t.Fatalf("GetUsers with query params writes to the cache, just like for all users")
	}
}

func TestBug14_C_POST(t *testing.T) {
	ctx, service, inputUser := Begin(nil)

	users_1, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	_, err = service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	users_2, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	postgresStub := service.userRepositoryPostgres.(*postgresRepoStub)

	if postgresStub.getUsersCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getUsersCalls)
	}

	if len(users_1) == len(users_2) {
		t.Fatalf("GetUsers with query params writes to the cache, just like for all users")
	}
}

func TestBug14_C_DEL(t *testing.T) {
	ctx, service, inputUser := Begin(nil)

	created, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	users_1, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	err = service.DeleteUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteUser() returned unexpected error: %v", err)
	}

	users_2, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	postgresStub := service.userRepositoryPostgres.(*postgresRepoStub)

	if postgresStub.getUsersCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getUsersCalls)
	}

	if len(users_1) == len(users_2) {
		t.Fatalf("GetUsers with query params writes to the cache, just like for all users")
	}
}

func TestBug14_C_PATCH(t *testing.T) {
	ctx, service, inputUser := Begin(nil)

	created, err := service.CreateUser(ctx, inputUser)
	if err != nil {
		t.Fatalf("CreateUser() returned unexpected error: %v", err)
	}

	users_1, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	user_1_fullname := users_1[len(users_1)-1].FullName

	full_name := "Ivan"
	_, err = service.PatchUser(ctx, created.ID, domain.UserPatch{
		FullName: domain.Nullable[string]{
			Set:   true,
			Value: &full_name,
		},
	})

	users_2, err := service.GetUsers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetUsers() returned unexpected error: %v", err)
	}

	users_2_fullname := users_2[len(users_2)-1].FullName

	postgresStub := service.userRepositoryPostgres.(*postgresRepoStub)

	if postgresStub.getUsersCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getUsersCalls)
	}

	if user_1_fullname == users_2_fullname {
		t.Fatalf("PatchUser with query params writes to the cache, just like for all users")
	}
}
