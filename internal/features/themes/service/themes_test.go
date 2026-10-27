package themes_service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

var (
	ErrRedisUnavailable       = errors.New("redis unavailable")
	ErrRedisConnectionRefused = errors.New("connection refused")
)

func generate_themes() []domain.Theme {
	return []domain.Theme{
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
		domain.NewThemeUnitialized("1234", nil, "Математика", 5000, "beginner", "3 месяца", "video", 1),
	}
}

func Begin(redisErr error) (context.Context, *ThemeService, domain.Theme) {
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

	postgres := &postgresRepoStub{themes: []domain.Theme{}}
	redis := &redisRepoStub{err: redisErr}

	service := &ThemeService{
		themeRepositoryPostgres: postgres,
		themeRepositoryRedis:    redis,
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
			cacheErr: ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputTheme := Begin(test.cacheErr)

			created, err := service.CreateTheme(ctx, inputTheme)
			if err != nil {
				t.Fatalf("CreateTheme() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(created, inputTheme) {
				t.Errorf("CreateTheme() = %v, want %v", created, inputTheme)
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
			cacheErr: ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputTheme := Begin(test.cacheErr)

			created, err := service.CreateTheme(ctx, inputTheme)
			if err != nil {
				t.Fatalf("CreateTheme() returned unexpected error: %v", err)
			}

			if err := service.DeleteTheme(ctx, created.ID); err != nil {
				t.Fatalf("DeleteTheme() returned unexpected error: %v", err)
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
			cacheErr: ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputTheme := Begin(test.cacheErr)

			created, err := service.CreateTheme(ctx, inputTheme)
			if err != nil {
				t.Fatalf("CreateTheme() returned unexpected error: %v", err)
			}

			got, err := service.GetTheme(ctx, inputTheme.ID)
			if err != nil {
				t.Fatalf("GetTheme() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, created) {
				t.Errorf("GetTheme() = %v, want %v", got, created)
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
			cacheErr: ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, _ := Begin(test.cacheErr)
			themes := generate_themes()

			for _, theme := range themes {
				if _, err := service.CreateTheme(ctx, theme); err != nil {
					t.Fatalf("CreateTheme() returned unexpected error: %v", err)
				}
			}

			got, err := service.GetThemes(ctx, nil, nil, nil)
			if err != nil {
				t.Fatalf("GetThemes() returned unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, themes) {
				t.Errorf("GetThemes() = %v, want %v", got, themes)
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
			cacheErr: ErrRedisConnectionRefused,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, service, inputTheme := Begin(test.cacheErr)

			created, err := service.CreateTheme(ctx, inputTheme)
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

			got, err := service.PatchTheme(ctx, created.ID, patch)
			if err != nil {
				t.Fatalf("PatchTheme() returned unexpected error: %v", err)
			}

			if got.Subject != subject {
				t.Errorf("PatchTheme() Subject = %q, want %q", got.Subject, subject)
			}
		})
	}
}

func TestPatchTheme_TwiceInRow(t *testing.T) {
	ctx, service, theme := Begin(ErrRedisUnavailable)

	created, err := service.CreateTheme(ctx, theme)
	if err != nil {
		t.Fatalf("CreateTheme() returned unexpected error: %v", err)
	}

	subject := "beginner"
	patch := domain.ThemePatch{
		Subject: domain.Nullable[string]{
			Set:   true,
			Value: &subject,
		},
	}

	theme_patch, err := service.PatchTheme(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchTheme() returned unexpected error: %v", err)
	}

	if theme_patch.Subject != subject {
		t.Errorf("PatchTheme() Subject = %q, want %q", theme_patch.Subject, subject)
	}

	format := "mixed"
	patch = domain.ThemePatch{
		Format: domain.Nullable[string]{
			Set:   true,
			Value: &format,
		},
	}

	theme_patch, err = service.PatchTheme(ctx, created.ID, patch)
	if err != nil {
		t.Fatalf("PatchTheme() returned unexpected error: %v", err)
	}

	if theme_patch.Format != format {
		t.Errorf("PatchTheme() Format = %q, want %q", theme_patch.Format, format)
	}
}
