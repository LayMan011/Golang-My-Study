package domain

import (
	"testing"
	"time"
)

func boolPtr(b bool) *bool {
	return &b
}

func TestThemeUserValidate(t *testing.T) {
	type testCase struct {
		name       string
		theme_user ThemeUser
		wantErr    bool
	}

	completedAt := time.Now()

	tests := []testCase{
		testCase{
			name: "CompletedAt == nil",
			theme_user: ThemeUser{
				Completed:   true,
				CompletedAt: nil,
			},
			wantErr: true,
		},
		testCase{
			name: "CompletedAt.Before(t.AdditionAt)",
			theme_user: ThemeUser{
				Completed:   true,
				AdditionAt:  time.Now().Add(1 * time.Second),
				CompletedAt: &completedAt,
			},
			wantErr: true,
		},
		testCase{
			name: "CompletedAt != nil",
			theme_user: ThemeUser{
				Completed:   false,
				CompletedAt: &completedAt,
			},
			wantErr: true,
		},
		testCase{
			name: "t.TotalLessons < t.CompletedLessons",
			theme_user: ThemeUser{
				Completed:        false,
				CompletedAt:      nil,
				TotalLessons:     0,
				CompletedLessons: 1,
			},
			wantErr: true,
		},
		testCase{
			name:       "return nil",
			theme_user: NewThemeUserUnitialized(1, 1),
			wantErr:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.theme_user.Validate()
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestThemeUserApplyPatch(t *testing.T) {
	type testCase struct {
		name    string
		patch   ThemeUserPatch
		wantErr bool
	}

	theme_user := NewThemeUserUnitialized(1, 1)

	tests := []testCase{
		testCase{
			name: "first validate",
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name: "Completed == true && Value == true",
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: boolPtr(true),
				},
			},
			wantErr: false,
		},
		testCase{
			name: "Completed == true && Value == false",
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: boolPtr(false),
				},
			},
			wantErr: false,
		},
		testCase{
			name: "return nil",
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   false,
					Value: nil,
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := theme_user.ApplyPatch(test.patch)
			if test.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
