package domain

import (
	"errors"
	"reflect"
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
	type testCase struct {
		name    string
		theme   Theme
		wantErr bool
	}

	tests := []testCase{
		testCase{
			name: "Title is empty",
			theme: Theme{
				Title:       "",
				Description: strPtr("0"),
				Subject:     strings.Repeat("λ", 100),
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "intermediate",
				Duration:    strings.Repeat("λ", 5),
				Format:      "mixed",
			},
			wantErr: true,
		},
		testCase{
			name: "LengthTitle = 101",
			theme: Theme{
				Title:       strings.Repeat("λ", 101),
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "LengthTitle = 100",
			theme: Theme{
				Title:       strings.Repeat("λ", 100),
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Length Description = 0",
			theme: Theme{
				Title:       "0",
				Description: new(string),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Length Description = 1001",
			theme: Theme{
				Title:       "0",
				Description: strPtr(strings.Repeat("λ", 1001)),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "advanced",
				Duration:    strings.Repeat("λ", 5),
				Format:      "text",
			},
			wantErr: true,
		},
		testCase{
			name: "LengthTitle = 100",
			theme: Theme{
				Title:       "0",
				Description: strPtr(strings.Repeat("λ", 1000)),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "advanced",
				Duration:    strings.Repeat("λ", 5),
				Format:      "text",
			},
			wantErr: false,
		},
		testCase{
			name: "Subject is empty",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "advanced",
				Duration:    strings.Repeat("λ", 5),
				Format:      "text",
			},
			wantErr: true,
		},
		testCase{
			name: "Length Subject = 101",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     strings.Repeat("λ", 101),
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "advanced",
				Duration:    strings.Repeat("λ", 5),
				Format:      "text",
			},
			wantErr: true,
		},
		testCase{
			name: "Length Subject = 100",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     strings.Repeat("λ", 100),
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "intermediate",
				Duration:    strings.Repeat("λ", 5),
				Format:      "mixed",
			},
			wantErr: false,
		},
		testCase{
			name: "Rating = -1",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(-1),
				Price:       100,
				Level:       "advanced",
				Duration:    strings.Repeat("λ", 5),
				Format:      "text",
			},
			wantErr: true,
		},
		testCase{
			name: "Rating = 0",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(0),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Rating = 5",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Rating = 6",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(6),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Price = -1",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       -1,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Price = 0",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       0,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Level is empty",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Duration is empty",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    "",
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Length Duration = 0",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    "0",
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Length Duration = 41",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("0", 41),
				Format:      "video",
			},
			wantErr: true,
		},
		testCase{
			name: "Length Duration = 0",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("0", 40),
				Format:      "video",
			},
			wantErr: false,
		},
		testCase{
			name: "Format is empty",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("0", 5),
				Format:      "",
			},
			wantErr: true,
		},
		testCase{
			name: "Full valid",
			theme: Theme{
				Title:       "0",
				Description: strPtr("0"),
				Subject:     "0",
				Rating:      float64Ptr(5),
				Price:       100,
				Level:       "beginner",
				Duration:    strings.Repeat("λ", 5),
				Format:      "video",
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.theme.Validate()
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if test.wantErr && !errors.Is(err, core_errors.ErrInvalidArgument) {
				t.Fatalf("expected ErrInvalidArgument, got %v", err)
			}
		})
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

func TestThemeValidateApplyPatch(t *testing.T) {
	type testCase struct {
		name        string
		theme       Theme
		patch       ThemePatch
		wantErr     bool
		wantChanged bool
	}

	description := "description"
	newTitle := "new title"

	tests := []testCase{
		testCase{
			name:  "invalid patch",
			theme: NewThemeUnitialized("title", &description, "00", 100, "beginner", "00", "video", 1),
			patch: ThemePatch{
				Title: Nullable[string]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name:  "all fields updated with invalid values",
			theme: NewThemeUnitialized("title", &description, "00", 100, "beginner", "00", "video", 1),
			patch: ThemePatch{
				Title:       Nullable[string]{Set: true, Value: &description},
				Description: Nullable[string]{Set: true, Value: &description},
				Subject:     Nullable[string]{Set: true, Value: &description},
				Level:       Nullable[string]{Set: true, Value: &description},
				Duration:    Nullable[string]{Set: true, Value: &description},
				Format:      Nullable[string]{Set: true, Value: &description},
			},
			wantErr: true,
		},
		testCase{
			name:  "update title",
			theme: NewThemeUnitialized("title", &description, "00", 100, "beginner", "00", "video", 1),
			patch: ThemePatch{
				Title: Nullable[string]{
					Set:   true,
					Value: &newTitle,
				},
			},
			wantChanged: true,
		},
		testCase{
			name:  "no fields updated",
			theme: NewThemeUnitialized("title", &description, "00", 100, "beginner", "00", "video", 1),
			patch: ThemePatch{
				Title:       Nullable[string]{Set: false},
				Description: Nullable[string]{Set: false},
				Subject:     Nullable[string]{Set: false},
				Level:       Nullable[string]{Set: false},
				Duration:    Nullable[string]{Set: false},
				Format:      Nullable[string]{Set: false},
			},
			wantChanged: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before := test.theme

			err := test.theme.ApplyPatch(test.patch)
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !reflect.DeepEqual(test.theme, before) {
					t.Errorf("theme changed after failed ApplyPatch(): got %#v, want unchanged %#v", test.theme, before)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			changed := !reflect.DeepEqual(before, test.theme)
			if changed != test.wantChanged {
				t.Errorf("theme changed = %v, want %v", changed, test.wantChanged)
			}
		})
	}
}
