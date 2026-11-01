package statistics_service

import (
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func TestCalcStatistics(t *testing.T) {
	type testCase struct {
		name        string
		themes_user []domain.ThemeUser

		wantCreated   int
		wantCompleted int
		wantRate      *float64
		wantAvgTime   *time.Duration
	}

	start := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	startAdd24 := start.Add(24 * time.Hour)
	startAdd72 := start.Add(72 * time.Hour)

	tests := []testCase{
		testCase{
			name:        "empty set",
			themes_user: nil,

			wantCreated:   0,
			wantCompleted: 0,
			wantRate:      nil,
			wantAvgTime:   nil,
		},
		testCase{
			name: "no completed themes_user",
			themes_user: []domain.ThemeUser{
				{AdditionAt: start},
				{AdditionAt: startAdd24},
			},

			wantCreated:   2,
			wantCompleted: 0,
			wantRate:      float64Ptr(0),
			wantAvgTime:   nil,
		},
		testCase{
			name: "all themes_user completed",
			themes_user: []domain.ThemeUser{
				{
					AdditionAt:  start,
					Completed:   true,
					CompletedAt: &startAdd24,
				},
				{
					AdditionAt:  start,
					Completed:   true,
					CompletedAt: &startAdd72,
				},
			},

			wantCreated:   2,
			wantCompleted: 2,
			wantRate:      float64Ptr(100),
			wantAvgTime:   durationPtr(48 * time.Hour),
		},
		testCase{
			name: "mixed completed and incomplete themes",
			themes_user: []domain.ThemeUser{
				domain.ThemeUser{
					AdditionAt:  start,
					Completed:   true,
					CompletedAt: &startAdd24,
				},
				domain.ThemeUser{AdditionAt: start},
				domain.ThemeUser{
					AdditionAt:  start,
					Completed:   true,
					CompletedAt: &startAdd72,
				},
				domain.ThemeUser{AdditionAt: start},
			},

			wantCreated:   4,
			wantCompleted: 2,
			wantRate:      float64Ptr(50),
			wantAvgTime:   durationPtr(48 * time.Hour),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calcStatistics(test.themes_user)

			if got.ThemesCreated != test.wantCreated {
				t.Errorf("ThemesCreated = %d, want %d", got.ThemesCreated, test.wantCreated)
			}

			if got.ThemesCompleted != test.wantCompleted {
				t.Errorf("ThemesCompleted = %d, want %d", got.ThemesCompleted, test.wantCompleted)
			}

			if !equalFloat64Ptr(got.ThemesCompletedRate, test.wantRate) {
				t.Errorf("ThemesCompletedRate = %v, want %v", got.ThemesCompletedRate, test.wantRate)
			}

			if !equalDurationPtr(got.ThemesAverageCompletionTime, test.wantAvgTime) {
				t.Errorf("ThemesAverageCompletionTime = %v, want %v", got.ThemesAverageCompletionTime, test.wantAvgTime)
			}
		})
	}
}

func float64Ptr(value float64) *float64 {
	return &value
}

func durationPtr(value time.Duration) *time.Duration {
	return &value
}

func equalFloat64Ptr(a, b *float64) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return *a == *b
}

func equalDurationPtr(a, b *time.Duration) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return *a == *b
}
