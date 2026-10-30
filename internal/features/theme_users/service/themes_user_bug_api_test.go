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

	redisErr := errors.New("redis unavailable")

	postgres := &postgresRepoStub{}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeUserService{
		themeUserRepositoryPostgres: postgres,
		themeUserRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestBug14_A(t *testing.T) {
	ctx, service, _ := Begin()

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
	ctx, service, _ := Begin()

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
	ctx, service, inputThemeUser := Begin()

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
	ctx, service, inputThemeUser := Begin()

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
	ctx, service, inputThemeUser := Begin()

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
