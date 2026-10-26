package themes_user_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func Begin() (context.Context, *ThemeUserService, domain.ThemeUser) {
	ctx := core_logger.ToContext(
		context.Background(),
		&core_logger.Logger{Logger: zap.NewNop()},
	)

	inputTheme := domain.ThemeUser{
		ID:               6,
		Version:          2,
		Completed:        false,
		AdditionAt:       time.Now(),
		CompletedAt:      nil,
		Percentages:      100,
		TotalLessons:     0,
		CompletedLessons: 0,
		ThemeID:          10,
		UserID:           3,
	}
	createdTheme := inputTheme

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{theme_user: createdTheme}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeUserService{
		themeUserRepositoryPostgres: postgres,
		themeUserRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestThemeService_CreateTheme(t *testing.T) {
	ctx, service, inputThemeUser := Begin()

	_, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}
}

func TestThemeService_DeleteTheme(t *testing.T) {
	ctx, service, inputThemeUser := Begin()

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	err = service.DeleteThemeUser(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteThemeUser() returned unexpected error: %v", err)
	}
}

func TestThemeService_GetTheme(t *testing.T) {
	ctx, service, inputThemeUser := Begin()

	_, err := service.GetThemeUser(ctx, inputThemeUser.ID)
	if err != nil {
		t.Fatalf("GetThemeUser() returned unexpected error: %v", err)
	}
}

func TestThemeService_GetThemes(t *testing.T) {
	ctx, service, _ := Begin()

	_, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}
}

func TestThemeService_PatchTheme(t *testing.T) {
	ctx, service, theme_user := Begin()

	_, err := service.CreateThemeUser(ctx, theme_user)
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

	_, err = service.PatchThemeUser(ctx, theme_user.ID, patch)
	if err != nil {
		t.Fatalf("PatchThemes() returned unexpected error: %v", err)
	}
}
