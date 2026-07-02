package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	themes, err := s.statisticsRepository.GetThemes(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get themes from repository: %w", err)
	}

	statistics := calcStatistics(themes)

	return statistics, nil
}

func calcStatistics(themes []domain.Theme) domain.Statistics {
	if len(themes) == 0 {
		return domain.NewStatistics(0, 0, nil, nil)
	}

	themesCreated := len(themes)

	themesCompleted := 0
	var totalCompletionDuration time.Duration
	for _, theme := range themes {
		if theme.Completed == true {
			themesCompleted++
		}

		completionDuration := theme.CompletionDuration()
		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	themesCompletedRate := float64(themesCompleted) / float64(themesCreated) * 100

	var themesAverageComplitionTime *time.Duration
	if themesCompleted > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(themesCompleted)

		themesAverageComplitionTime = &avg
	}

	return domain.NewStatistics(
		themesCreated,
		themesCompleted,
		&themesCompletedRate,
		themesAverageComplitionTime,
	)
}
