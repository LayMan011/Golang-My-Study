package themes_service

import (
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin(nil)

	themes_generate := generate_themes()
	for _, theme := range themes_generate {
		if _, err := service.CreateTheme(ctx, theme); err != nil {
			t.Fatalf("CreateTheme() returned unexpected error: %v", err)
		}
	}

	user_id := 2
	themes_query, err := service.GetThemes(ctx, &user_id, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() with query params returned unexpected error: %v", err)
	}

	if len(themes_query) != 2 {
		t.Fatalf("expected 2 themes records for user 2, got %d", len(themes_query))
	}

	if themes_query[0].ID == themes_query[1].ID {
		t.Fatalf("expected distinct theme IDs, got %d for both records", themes_query[0].ID)
	}

	if themes_query[0].AuthorUserID != user_id {
		t.Fatalf("theme ID %d: expected AuthorUserID %d, got %d", themes_query[0].ID, user_id, themes_query[0].AuthorUserID)
	}

	if themes_query[1].AuthorUserID != user_id {
		t.Fatalf("theme ID %d: expected AuthorUserID %d, got %d", themes_query[1].ID, user_id, themes_query[1].AuthorUserID)
	}

	themes, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes_query) == len(themes) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_B(t *testing.T) {
	ctx, service, _ := Begin(nil)

	themes_generate := generate_themes()
	for _, theme := range themes_generate {
		if _, err := service.CreateTheme(ctx, theme); err != nil {
			t.Fatalf("CreateTheme() returned unexpected error: %v", err)
		}
	}

	limit := 0
	themes_query, err := service.GetThemes(ctx, nil, &limit, nil)
	if err != nil {
		t.Fatalf("GetThemes() with query params returned unexpected error: %v", err)
	}

	if len(themes_query) != 0 {
		t.Fatalf("expected no themes with limit %d, got %d", limit, len(themes_query))
	}

	themes, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes) != len(themes_generate) {
		t.Fatalf("expected %d themes without filters, got %d", len(themes_generate), len(themes))
	}

	if len(themes_query) == len(themes) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_POST(t *testing.T) {
	ctx, service, inputTheme := Begin(nil)

	postgresStub := service.themeRepositoryPostgres.(*postgresRepoStub)

	_, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	_, err = service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_2, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	_, err = service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	if len(themes_1) == len(themes_2) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_DEL(t *testing.T) {
	ctx, service, inputTheme := Begin(nil)

	postgresStub := service.themeRepositoryPostgres.(*postgresRepoStub)

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	_, err = service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	err = service.DeleteTheme(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteTheme() returned unexpected error: %v", err)
	}

	themes_2, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	_, err = service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	if len(themes_1) == len(themes_2) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_PATCH(t *testing.T) {
	ctx, service, inputTheme := Begin(nil)

	postgresStub := service.themeRepositoryPostgres.(*postgresRepoStub)

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	_, err = service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 1 {
		t.Fatalf("expected 1 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	theme_1_title := themes_1[len(themes_1)-1].Title

	title := "1234567"
	_, err = service.PatchTheme(ctx, created.ID, domain.ThemePatch{
		Title: domain.Nullable[string]{
			Set:   true,
			Value: &title,
		},
	})

	themes_2, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	_, err = service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if postgresStub.getThemesCalls != 2 {
		t.Fatalf("expected 2 Postgres list read, got %d", postgresStub.getThemesCalls)
	}

	theme_2_title := themes_2[len(themes_2)-1].Title

	if theme_1_title == theme_2_title {
		t.Fatalf("PatchTheme with query params writes to the cache, just like for all themes")
	}
}
