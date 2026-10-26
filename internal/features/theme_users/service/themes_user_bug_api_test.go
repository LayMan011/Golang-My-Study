package themes_user_service

import (
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin(core_errors.ErrRedisUnavailable)

	user_id := 1
	themes_user_query, err := service.GetThemesUser(ctx, &user_id, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() with query params returned unexpected error: %v", err)
	}

	themes_user, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user_query) != len(themes_user) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_B(t *testing.T) {
	ctx, service, _ := Begin(core_errors.ErrRedisUnavailable)

	limit := 0
	themes_user_query, err := service.GetThemesUser(ctx, nil, nil, &limit, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() with query params returned unexpected error: %v", err)
	}

	themes_user, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user_query) != len(themes_user) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_POST(t *testing.T) {
	ctx, service, inputThemeUser := Begin(core_errors.ErrRedisUnavailable)

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	_, err = service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_user_2, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user_1) == len(themes_user_2) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_DEL(t *testing.T) {
	ctx, service, inputThemeUser := Begin(core_errors.ErrRedisUnavailable)

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	err = service.DeleteThemeUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteThemeUser() returned unexpected error: %v", err)
	}

	themes_user_2, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if len(themes_user_1) == len(themes_user_2) {
		t.Fatalf("GetThemesUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug14_C_PATCH(t *testing.T) {
	ctx, service, inputThemeUser := Begin(core_errors.ErrRedisUnavailable)

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	themes_user_1, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
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

	theme_user_2_completed := themes_user_2[len(themes_user_2)-1].Completed

	if theme_user_1_completed == theme_user_2_completed {
		t.Fatalf("PatchThemeUser with query params writes to the cache, just like for all themes_user")
	}
}

func TestBug23(t *testing.T) {
	ctx, service, theme_user := Begin(core_errors.ErrRedisUnavailable)

	created, err := service.CreateThemeUser(ctx, theme_user)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	completed := false
	patch := domain.ThemeUserPatch{
		Completed: domain.Nullable[bool]{
			Set:   true,
			Value: &completed,
		},
	}

	theme_user_patch, err := service.PatchThemeUser(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchTheme() returned unexpected error: %v", err)
	}

	if theme_user_patch.Completed != completed {
		t.Errorf("PatchThemeUser() Completed = %t, want %t", theme_user_patch.Completed, completed)
	}

	completed = true
	patch = domain.ThemeUserPatch{
		Completed: domain.Nullable[bool]{
			Set:   true,
			Value: &completed,
		},
	}

	theme_user_patch, err = service.PatchThemeUser(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchTheme() returned unexpected error: %v", err)
	}

	if theme_user_patch.Completed != completed {
		t.Errorf("PatchThemeUser() Completed = %t, want %t", theme_user_patch.Completed, completed)
	}
}
