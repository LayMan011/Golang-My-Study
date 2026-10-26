package themes_user_service

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

func generate_themes_user() []domain.ThemeUser {
	return []domain.ThemeUser{
		domain.NewThemeUserUnitialized(1, 1),
		domain.NewThemeUserUnitialized(1, 1),
		domain.NewThemeUserUnitialized(1, 1),
	}
}

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

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{themes_user: []domain.ThemeUser{}}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeUserService{
		themeUserRepositoryPostgres: postgres,
		themeUserRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestThemeService_CreateTheme(t *testing.T) {
	ctx, service, inputThemeUser := Begin()

	empty_theme_user := domain.ThemeUser{}

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(created, empty_theme_user) {
		t.Errorf("CreateThemeUser() = %v, want %v", created, empty_theme_user)
	}

	if !reflect.DeepEqual(created, inputThemeUser) {
		t.Errorf("CreateThemeUser() = %v, want %v", created, inputThemeUser)
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

	empty_theme_user := domain.ThemeUser{}

	created, err := service.CreateThemeUser(ctx, inputThemeUser)
	if err != nil {
		t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
	}

	get_theme_user, err := service.GetThemeUser(ctx, inputThemeUser.ID)
	if err != nil {
		t.Fatalf("GetThemeUser() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_theme_user, empty_theme_user) {
		t.Errorf("GetThemeUser() = %v, want %v", get_theme_user, empty_theme_user)
	}

	if !reflect.DeepEqual(get_theme_user, created) {
		t.Errorf("GetThemeUser() = %v, want %v", get_theme_user, created)
	}
}

func TestThemeService_GetThemes(t *testing.T) {
	ctx, service, _ := Begin()

	empty_themes_user := []domain.ThemeUser{}

	themes_user := generate_themes_user()

	for _, theme_user := range themes_user {
		_, err := service.CreateThemeUser(ctx, theme_user)
		if err != nil {
			t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
		}
	}

	get_themes_user, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
	}

	if reflect.DeepEqual(get_themes_user, empty_themes_user) {
		t.Errorf("GetThemeUsers() = %v, want %v", get_themes_user, empty_themes_user)
	}

	if !reflect.DeepEqual(get_themes_user, themes_user) {
		t.Errorf("GetThemeUsers() = %v, want %v", get_themes_user, themes_user)
	}
}

func TestThemeService_PatchTheme(t *testing.T) {
	ctx, service, theme_user := Begin()

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
}
