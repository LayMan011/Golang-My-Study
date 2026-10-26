package domain

import (
	"reflect"
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
				Completed:        true,
				CompletedAt:      nil,
				TotalLessons:     1,
				CompletedLessons: 1,
			},
			wantErr: true,
		},
		testCase{
			name: "CompletedAt.Before(t.AdditionAt)",
			theme_user: ThemeUser{
				Completed:        true,
				AdditionAt:       time.Now().Add(1 * time.Second),
				CompletedAt:      &completedAt,
				TotalLessons:     1,
				CompletedLessons: 1,
			},
			wantErr: true,
		},
		testCase{
			name: "CompletedAt != nil",
			theme_user: ThemeUser{
				Completed:        false,
				CompletedAt:      &completedAt,
				TotalLessons:     1,
				CompletedLessons: 1,
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
		name       string
		theme_user ThemeUser
		patch      ThemeUserPatch
		wantErr    bool
	}

	fullPercentages := 100
	completedAt := time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)
	userWithProgress := NewThemeUserUnitialized(1, 1)
	userWithProgress.Completed = true
	userWithProgress.Percentages = 75
	userWithProgress.CompletedAt = &completedAt

	tests := []testCase{
		testCase{
			name:       "first validate",
			theme_user: NewThemeUserUnitialized(1, 1),
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: nil,
				},
			},
			wantErr: true,
		},
		testCase{
			name:       "Completed == true && Value == true",
			theme_user: NewThemeUserUnitialized(1, 1),
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: boolPtr(true),
				},
			},
			wantErr: false,
		},
		testCase{
			name:       "Completed == true && Value == false",
			theme_user: userWithProgress,
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{
					Set:   true,
					Value: boolPtr(false),
				},
			},
			wantErr: false,
		},
		testCase{
			name:       "no fields updated",
			theme_user: NewThemeUserUnitialized(1, 1),
			patch: ThemeUserPatch{
				Completed: Nullable[bool]{Set: false, Value: nil},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before := test.theme_user

			percentagesBefore := test.theme_user.Percentages
			completedAtBefore := test.theme_user.CompletedAt

			err := test.theme_user.ApplyPatch(test.patch)
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !reflect.DeepEqual(test.theme_user, before) {
					t.Errorf("theme_user changed after failed ApplyPatch(): got %#v, want unchanged %#v", test.theme_user, before)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if test.patch.Completed.Set {
				if test.theme_user.Completed != *test.patch.Completed.Value {
					t.Errorf("Completed = %v, want %v",
						test.theme_user.Completed, *test.patch.Completed.Value)
				}

				if test.patch.Completed.Value != nil && !*test.patch.Completed.Value {
					if test.theme_user.Percentages != 0 {
						t.Errorf("Percentages = %v, want 0", test.theme_user.Percentages)
					}
					if test.theme_user.CompletedAt != nil {
						t.Errorf("CompletedAt = %v, want nil", test.theme_user.CompletedAt)
					}
				} else if test.patch.Completed.Value != nil && *test.patch.Completed.Value {
					if test.theme_user.Percentages != fullPercentages {
						t.Errorf("Percentages = %v, want %d", test.theme_user.Percentages, fullPercentages)
					}
				}
			} else {
				if test.theme_user.Percentages != percentagesBefore {
					t.Errorf("Percentages = %v, want unchanged value %v",
						test.theme_user.Percentages, percentagesBefore)
				}
				if test.theme_user.CompletedAt != completedAtBefore {
					t.Errorf("CompletedAt = %v, want unchanged value %v",
						test.theme_user.CompletedAt, completedAtBefore)
				}
			}
		})
	}
}
