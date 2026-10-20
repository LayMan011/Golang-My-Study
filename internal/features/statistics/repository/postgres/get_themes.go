package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *StatisticsRepository) GetThemeUsers(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) ([]domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT id, version, completed, addition_at, completed_at, percentages, theme_id, user_id, total_lessons, completed_lessons
	FROM progress.themes_user
	`)

	args := []any{}
	conditions := []string{}

	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("user_id=$%d", len(args)+1))
		args = append(args, userID)
	}

	if from != nil {
		conditions = append(conditions, fmt.Sprintf("addition_at>=$%d", len(args)+1))
		args = append(args, from)
	}

	if to != nil {
		conditions = append(conditions, fmt.Sprintf("addition_at<$%d", len(args)+1))
		args = append(args, to)
	}

	if len(conditions) > 0 {
		str := " WHERE " + strings.Join(conditions, " AND ")
		queryBuilder.WriteString(str)
	}

	queryBuilder.WriteString(" ORDER BY id ASC;")
	rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("select themes: %w", err)
	}
	defer rows.Close()

	var themeModels []ThemeUserModel
	for rows.Next() {
		var themeModel ThemeUserModel

		err := rows.Scan(
			&themeModel.ID,
			&themeModel.Version,
			&themeModel.Completed,
			&themeModel.AdditionAt,
			&themeModel.CompletedAt,
			&themeModel.Percentages,
			&themeModel.ThemeID,
			&themeModel.UserID,
			&themeModel.TotalLessons,
			&themeModel.CompletedLessons,
		)
		if err != nil {
			return nil, fmt.Errorf("scan themes: %w", err)
		}

		themeModels = append(themeModels, themeModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	themeDomains := themesUserDomainFromModels(themeModels)

	return themeDomains, nil
}
