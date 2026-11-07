package themes_user_service

import (
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin(nil)

	themesUser_generate := generate_themes_user()
	for _, themeUser := range themesUser_generate {
		if _, err := service.CreateThemeUser(ctx, themeUser); err != nil {
			t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
		}
	}

	user_id := 2
	themes_user_query, err := service.GetThemesUser(ctx, &user_id, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() with query params returned unexpected error: %v", err)
	}

	if len(themes_user_query) != 2 {
		t.Fatalf("expected 2 theme_user records for user %d, got %d", user_id, len(themes_user_query))
	}

	if themes_user_query[0].ID == themes_user_query[1].ID {
		t.Fatalf("expected distinct theme_user IDs, got %d for both records", themes_user_query[0].ID)
	}

	if themes_user_query[0].UserID != user_id {
		t.Fatalf("theme_user ID %d: expected UserID %d, got %d", themes_user_query[0].ID, user_id, themes_user_query[0].UserID)
	}

	if themes_user_query[1].UserID != user_id {
		t.Fatalf("theme_user ID %d: expected UserID %d, got %d", themes_user_query[1].ID, user_id, themes_user_query[1].UserID)
	}

	themes_user, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user_query) == len(themes_user) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_B(t *testing.T) {
	ctx, service, _ := Begin(nil)

	themesUser_generate := generate_themes_user()
	for _, themeUser := range themesUser_generate {
		if _, err := service.CreateThemeUser(ctx, themeUser); err != nil {
			t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
		}
	}

	limit := 0
	themes_user_query, err := service.GetThemesUser(ctx, nil, nil, &limit, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() with query params returned unexpected error: %v", err)
	}

	if len(themes_user_query) != 0 {
		t.Fatalf("expected no theme_user records with limit %d, got %d", limit, len(themes_user_query))
	}

	themes_user, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user) != len(themesUser_generate) {
		t.Fatalf("expected %d theme_user records without filters, got %d", len(themesUser_generate), len(themes_user))
	}

	if len(themes_user_query) == len(themes_user) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_POST(t *testing.T) {
	ctx, service, inputThemeUser := Begin(nil)

	postgresStub := service.themeUserRepositoryPostgres.(*postgresRepoStub)

	_, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	_, err = service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_user_2, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	_, err = service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	if len(themes_user_1) == len(themes_user_2) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_DEL(t *testing.T) {
	ctx, service, inputThemeUser := Begin(nil)

	postgresStub := service.themeUserRepositoryPostgres.(*postgresRepoStub)

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	_, err = service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	err = service.DeleteThemeUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteThemeUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	themes_user_2, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	_, err = service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	if len(themes_user_1) == len(themes_user_2) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_PATCH(t *testing.T) {
	ctx, service, inputThemeUser := Begin(nil)

	postgresStub := service.themeUserRepositoryPostgres.(*postgresRepoStub)

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	_, err = service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	theme_user_1_completed := themes_user_1[len(themes_user_1)-1].Completed

	completed := true
	_, err = service.PatchThemeUser(ctx, created.ID, domain.ThemeUserPatch{
		Completed: domain.Nullable[bool]{
			Set:   true,
			Value: &completed,
		},
	})

	themes_user_2, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	_, err = service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesUserCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesUserCalls)
	}

	theme_user_2_completed := themes_user_2[len(themes_user_2)-1].Completed

	if theme_user_1_completed == theme_user_2_completed {
		t.Fatalf("PatchThemeUser with query params writes to the cache, just like for all themes_user")
	}
}
