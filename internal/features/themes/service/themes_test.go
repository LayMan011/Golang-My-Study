package themes_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

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
	createdTheme := inputTheme

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{theme: createdTheme}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeService{
		hash_once:               make(map[string]bool),
		hash_many:               make(map[string]bool),
		themeRepositoryPostgres: postgres,
		themeRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestThemeService_CreateTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	_, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}
}

func TestThemeService_DeleteTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	_, err := service.CreateTheme(ctx, inputTheme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	err = service.DeleteTheme(ctx, inputTheme.ID)
	if err != nil {
		t.Fatalf("DeleteTheme() returned unexpected error: %v", err)
	}
}

func TestThemeService_GetTheme(t *testing.T) {
	ctx, service, inputTheme := Begin()

	_, err := service.GetTheme(ctx, inputTheme.ID)
	if err != nil {
		t.Fatalf("GetTheme() returned unexpected error: %v", err)
	}
}

func TestThemeService_GetThemes(t *testing.T) {
	ctx, service, _ := Begin()

	_, err := service.GetThemes(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemes() returned unexpected error: %v", err)
	}
}

func TestThemeService_PatchTheme(t *testing.T) {
	ctx, service, theme := Begin()

	_, err := service.CreateTheme(ctx, theme)
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

	_, err = service.PatchTheme(ctx, theme.ID, patch)
	if err != nil {
		t.Fatalf("PatchThemes() returned unexpected error: %v", err)
	}
}
