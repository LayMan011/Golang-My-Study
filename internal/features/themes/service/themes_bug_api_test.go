package themes_service

import (
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin(ErrRedisUnavailable)

	user_id := 1
	themes_query, err := service.GetThemes(ctx, &user_id, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() with query params returned unexpected error: %v", err)
	}

	themes, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes_query) != len(themes) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_B(t *testing.T) {
	ctx, service, _ := Begin(ErrRedisUnavailable)

	limit := 0
	themes_query, err := service.GetThemes(ctx, nil, &limit, nil)
	if err != nil {
		t.Fatalf("GetThemes() with query params returned unexpected error: %v", err)
	}

	themes, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes_query) != len(themes) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_POST(t *testing.T) {
	ctx, service, inputTheme := Begin(ErrRedisUnavailable)

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	_, err = service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_2, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes_1) == len(themes_2) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_DEL(t *testing.T) {
	ctx, service, inputTheme := Begin(ErrRedisUnavailable)

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	err = service.DeleteTheme(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteTheme() returned unexpected error: %v", err)
	}

	themes_2, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if len(themes_1) == len(themes_2) {
		t.Fatalf("GetThemes with query params writes to the cache, just like for all themes")
	}
}

func TestBug14_C_PATCH(t *testing.T) {
	ctx, service, inputTheme := Begin(ErrRedisUnavailable)

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	themes_1, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
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

	theme_2_title := themes_2[len(themes_2)-1].Title

	if theme_1_title == theme_2_title {
		t.Fatalf("PatchTheme with query params writes to the cache, just like for all themes")
	}
}
