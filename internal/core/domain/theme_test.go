package domain

import (
	"errors"
	"strings"
	"testing"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func strPtr(s string) *string {
	return &s
}

func float64Ptr(s float64) *float64 {
	return &s
}

func TestValidateTheme(t *testing.T) {
	themes := []Theme{
		Theme{
			Title: "",
		},
		Theme{
			Title: strings.Repeat("0", 101),
		},
		Theme{
			Title:       "0",
			Description: nil,
		},
		Theme{
			Title:       "0",
			Description: new(string),
		},
		Theme{
			Title:       "0",
			Description: strPtr(strings.Repeat("0", 1001)),
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "",
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     strings.Repeat("0", 101),
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      nil,
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      new(float64),
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(6),
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       -1,
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       100,
			Level:       "",
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       100,
			Level:       "beginner",
			Duration:    "",
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       100,
			Level:       "beginner",
			Duration:    strings.Repeat("0", 50),
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       100,
			Level:       "beginner",
			Duration:    strings.Repeat("0", 5),
			Format:      "",
		},
		Theme{
			Title:       "0",
			Description: strPtr("0"),
			Subject:     "0",
			Rating:      float64Ptr(5),
			Price:       100,
			Level:       "beginner",
			Duration:    strings.Repeat("0", 5),
			Format:      "video",
		},
	}

	for _, theme := range themes {
		theme.Validate()
	}
}

func TestValidateThemePatch(t *testing.T) {
	type testCase struct {
		name    string
		patch   ThemePatch
		wantErr bool
	}

	tests := []testCase{
		// Title = nil
		{
			name: "Title nil",
			patch: ThemePatch{
				Title: Nullable[string]{Set: true, Value: nil},
			},
			wantErr: true,
		},
		// Subject = nil (Title валиден, Subject Set=false → пропускается)
		{
			name: "Subject not set",
			patch: ThemePatch{
				Title:   Nullable[string]{Set: true, Value: strPtr("0")},
				Subject: Nullable[string]{Set: false, Value: nil},
			},
			wantErr: false,
		},
		// Subject = nil, Set=true
		{
			name: "Subject nil",
			patch: ThemePatch{
				Title:   Nullable[string]{Set: true, Value: strPtr("0")},
				Subject: Nullable[string]{Set: true, Value: nil},
			},
			wantErr: true,
		},
		// Level = nil
		{
			name: "Level nil",
			patch: ThemePatch{
				Title:   Nullable[string]{Set: true, Value: strPtr("0")},
				Subject: Nullable[string]{Set: true, Value: strPtr("0")},
				Level:   Nullable[string]{Set: true, Value: nil},
			},
			wantErr: true,
		},
		// Duration = nil
		{
			name: "Duration nil",
			patch: ThemePatch{
				Title:    Nullable[string]{Set: true, Value: strPtr("0")},
				Subject:  Nullable[string]{Set: true, Value: strPtr("0")},
				Level:    Nullable[string]{Set: true, Value: strPtr("0")},
				Duration: Nullable[string]{Set: true, Value: nil},
			},
			wantErr: true,
		},
		// Format = nil
		{
			name: "Format nil",
			patch: ThemePatch{
				Title:    Nullable[string]{Set: true, Value: strPtr("0")},
				Subject:  Nullable[string]{Set: true, Value: strPtr("0")},
				Level:    Nullable[string]{Set: true, Value: strPtr("0")},
				Duration: Nullable[string]{Set: true, Value: strPtr("0")},
				Format:   Nullable[string]{Set: true, Value: nil},
			},
			wantErr: true,
		},
		// Всё валидно → nil
		{
			name: "all valid",
			patch: ThemePatch{
				Title:    Nullable[string]{Set: true, Value: strPtr("0")},
				Subject:  Nullable[string]{Set: true, Value: strPtr("0")},
				Level:    Nullable[string]{Set: true, Value: strPtr("0")},
				Duration: Nullable[string]{Set: true, Value: strPtr("0")},
				Format:   Nullable[string]{Set: true, Value: strPtr("0")},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.patch.Validate()
			if tc.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if tc.wantErr && !errors.Is(err, core_errors.ErrInvalidArgument) {
				t.Fatalf("expected ErrInvalidArgument, got %v", err)
			}
		})
	}
}

func TestValidateApplyPatch(t *testing.T) {
	type testCase struct {
		name    string
		patch   ThemePatch
		wantErr bool
	}

	description := "description"
	theme := NewThemeUnitialized("title", &description, "00", 100, "beginner", "00", "video", 1)

	tests := []testCase{
		testCase{
			name: "non validate",
			patch: ThemePatch{
				Title: Nullable[string]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name: "all rewies",
			patch: ThemePatch{
				Title: Nullable[string]{
					Set:   true,
					Value: &description,
				},
				Description: Nullable[string]{
					Set:   true,
					Value: &description,
				},
				Subject: Nullable[string]{
					Set:   true,
					Value: &description,
				},
				Level: Nullable[string]{
					Set:   true,
					Value: &description,
				},
				Duration: Nullable[string]{
					Set:   true,
					Value: &description,
				},
				Format: Nullable[string]{
					Set:   true,
					Value: &description,
				},
			},
			wantErr: true,
		},
		testCase{
			name: "return nil",
			patch: ThemePatch{
				Title: Nullable[string]{
					Set:   false,
					Value: nil,
				},
				Description: Nullable[string]{
					Set:   false,
					Value: nil,
				},
				Subject: Nullable[string]{
					Set:   false,
					Value: nil,
				},
				Level: Nullable[string]{
					Set:   false,
					Value: nil,
				},
				Duration: Nullable[string]{
					Set:   false,
					Value: nil,
				},
				Format: Nullable[string]{
					Set:   false,
					Value: nil,
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := theme.ApplyPatch(test.patch)
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestThemePatch(t *testing.T) {
	NewThemePatch(
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
		Nullable[string]{
			Set:   false,
			Value: nil,
		},
	)
}
