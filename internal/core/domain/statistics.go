package domain

import "time"

type Statistics struct {
	ThemesCreated               int
	ThemesCompleted             int
	ThemesCompletedRate         *float64
	ThemesAverageCompletionTime *time.Duration
}

func NewStatistics(
	themesCreated int,
	themesCompleted int,
	themesCompletedRate *float64,
	themeAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		ThemesCreated:               themesCreated,
		ThemesCompleted:             themesCompleted,
		ThemesCompletedRate:         themesCompletedRate,
		ThemesAverageCompletionTime: themeAverageCompletionTime,
	}
}
