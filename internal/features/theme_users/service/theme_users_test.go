package themes_user_service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
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

func Begin(redisErr error) (context.Context, *ThemeUserService, domain.ThemeUser) {
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

	postgres := &postgresRepoStub{themes_user: []domain.ThemeUser{}}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeUserService{
		themeUserRepositoryPostgres: postgres,
		themeUserRepositoryRedis:    redis,
	}

	return ctx, service, inputTheme
}

func TestThemeService_CreateTheme(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputThemeUser := Begin(test.cacheErr)

			created, err := service.CreateThemeUser(ctx, inputThemeUser)
			if err != nil {
				t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(created, inputThemeUser) {
				t.Errorf("CreateThemeUser() = %v, want %v", created, inputThemeUser)
			}
		})
	}
}

func TestThemeService_DeleteTheme(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputThemeUser := Begin(test.cacheErr)

			created, err := service.CreateThemeUser(ctx, inputThemeUser)
			if err != nil {
				t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
			}

			if err := service.DeleteThemeUser(ctx, created.ID); err != nil {
				t.Fatalf("DeleteThemeUser() returned unexpected error: %v", err)
			}
		})
	}
}

func TestThemeService_GetTheme(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputThemeUser := Begin(test.cacheErr)

			created, err := service.CreateThemeUser(ctx, inputThemeUser)
			if err != nil {
				t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
			}

			got, err := service.GetThemeUser(ctx, inputThemeUser.ID)
			if err != nil {
				t.Fatalf("GetThemeUser() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, created) {
				t.Errorf("GetThemeUser() = %v, want %v", got, created)
			}
		})
	}
}

func TestThemeService_GetThemes(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, _ := Begin(test.cacheErr)
			themesUser := generate_themes_user()

			for _, themeUser := range themesUser {
				if _, err := service.CreateThemeUser(ctx, themeUser); err != nil {
					t.Fatalf("CreateThemeUser() returned unexpected error: %v", err)
				}
			}

			got, err := service.GetThemesUser(ctx, nil, nil, nil, nil)
			if err != nil {
				t.Fatalf("GetThemesUser() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, themesUser) {
				t.Errorf("GetThemesUser() = %v, want %v", got, themesUser)
			}
		})
	}
}

func TestThemeService_PatchTheme(t *testing.T) {
	type testCase struct {
		name     string
		cacheErr error
	}

	tests := []testCase{
		testCase{
			name:     "cache miss",
			cacheErr: core_errors.ErrNotFound,
		},
		testCase{
			name:     "redis error",
			cacheErr: core_errors.ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputThemeUser := Begin(test.cacheErr)

			created, err := service.CreateThemeUser(ctx, inputThemeUser)
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

			got, err := service.PatchThemeUser(ctx, created.ID, patch)
			if err != nil {
				t.Fatalf("PatchThemeUser() returned unexpected error: %v", err)
			}

			if got.Completed != completed {
				t.Errorf("PatchThemeUser() Completed = %t, want %t", got.Completed, completed)
			}
		})
	}
}
