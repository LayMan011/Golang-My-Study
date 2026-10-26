package themes_service

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

func generate_themes() []domain.Theme {
	return []domain.Theme{
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
	}
}

func Begin() (context.Context, *ThemeService, domain.Theme) {
	ctx := core_logger.ToContext(
		context.Background(),
		&core_logger.Logger{Logger: zap.NewNop()},
	)

	inputTheme := domain.Theme{
		ID:              21,
		Version:         1,
		Title:           "123",
		Description:     nil,
		CreatedAt:       time.Now(),
		Subject:         "Русский язык",
		Rating:          nil,
		AllRatings:      0,
		NumberOfRatings: 0,
		NumberOfUsers:   0,
		Price:           5000,
		Level:           "beginner",
		Duration:        "3 месяца",
		Format:          "video",
		AuthorUserID:    26,
	}

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{themes: []domain.Theme{}}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeService{
		themeRepositoryPostgres: postgres,
		themeRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestThemeService_CreateTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	empty_theme := domain.Theme{}

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(created, empty_theme) {
		t.Errorf("CreateTheme() = %v, want %v", created, empty_theme)
	}

	if !reflect.DeepEqual(created, inputTheme) {
		t.Errorf("CreateTheme() = %v, want %v", created, inputTheme)
	}
}

func TestThemeService_DeleteTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	err = service.DeleteTheme(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteTheme() returned unexpected error: %v", err)
	}
}

func TestThemeService_GetTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	empty_theme := domain.Theme{}

	created, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	get_theme, err := service.GetTheme(ctx, inputTheme.ID)
	if err != nil {
		t.Fatalf("GetTheme() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_theme, empty_theme) {
		t.Errorf("GetTheme() = %v, want %v", get_theme, empty_theme)
	}

	if !reflect.DeepEqual(get_theme, created) {
		t.Errorf("GetTheme() = %v, want %v", get_theme, created)
	}
}

func TestThemeService_GetThemes(t *testing.T) {
	ctx, service, _ := Begin()

	empty_themes := []domain.ThemeUser{}

	themes := generate_themes()

	for _, theme := range themes {
		_, err := service.CreateTheme(ctx, theme)
		if err != nil {
			t.Fatalf("CreateTheme() returned unexpected error: %v", err)
		}
	}

	get_themes, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_themes, empty_themes) {
		t.Errorf("GetThemes() = %v, want %v", get_themes, empty_themes)
	}

	if !reflect.DeepEqual(get_themes, themes) {
		t.Errorf("GetThemes() = %v, want %v", get_themes, themes)
	}
}

func TestThemeService_PatchTheme(t *testing.T) {
	ctx, service, theme := Begin()

	created, err := service.CreateTheme(ctx, theme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	subject := "beginner"
	patch := domain.ThemePatch{
		Title: domain.Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Description: domain.Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Subject: domain.Nullable[string]{
			Set:   true,
			Value: &subject,
		},
		Level: domain.Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Duration: domain.Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Format: domain.Nullable[string]{
			Set:   false,
			Value: nil,
		},
	}

	theme_patch, err := service.PatchTheme(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchTheme() returned unexpected error: %v", err)
	}

	if theme_patch.Subject != subject {
		t.Errorf("PatchTheme() Subject = %q, want %q", theme_patch.Subject, subject)
	}
}
