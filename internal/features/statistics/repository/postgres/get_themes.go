package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *StatisticsRepository) GetThemes(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) ([]domain.Theme, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, percentages, author_user_id
	FROM progress.themes
	`)

	args := []any{}
	conditions := []string{}

	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)+1))
		args = append(args, userID)
	}

	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}

	if to != nil {
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)+1))
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

	var themeModels []ThemeModel
	for rows.Next() {
		var themeModel ThemeModel

		err := rows.Scan(
			&themeModel.ID,
			&themeModel.Version,
			&themeModel.Title,
			&themeModel.Description,
			&themeModel.Completed,
			&themeModel.CreatedAt,
			&themeModel.CompletedAt,
			&themeModel.Percentages,
			&themeModel.AuthorUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan themes: %w", err)
		}

		themeModels = append(themeModels, themeModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	themeDomains := themesDomainFromModels(themeModels)

	return themeDomains, nil
}
