package themes_user_postgres_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeUserRepository) GetThemesUser(
	ctx context.Context,
	userID *int,
	themeID *int,
	first *int,
	last *int,
) ([]domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, completed, addition_at, completed_at, percentages, total_lessons, completed_lessons, theme_id, user_id
	FROM progress.themes_user
	%s
	ORDER by id ASC
	LIMIT $1
	OFFSET $2;
	`

	args := []any{first, last}

	if userID != nil && themeID != nil {
		query = fmt.Sprintf(query, "WHERE user_id=$3 AND theme_id=$4")
		args = append(args, userID)
		args = append(args, themeID)
	} else if userID != nil && themeID == nil {
		query = fmt.Sprintf(query, "WHERE user_id=$3")
		args = append(args, userID)
	} else if userID == nil && themeID != nil {
		query = fmt.Sprintf(query, "WHERE theme_id=$3")
		args = append(args, themeID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select themes: %w", err)
	}
	defer rows.Close()

	var themeUserModels []ThemeUserModel
	for rows.Next() {
		var themeUserModel ThemeUserModel

		err := rows.Scan(
			&themeUserModel.ID,
			&themeUserModel.Version,
			&themeUserModel.Completed,
			&themeUserModel.AdditionAt,
			&themeUserModel.CompletedAt,
			&themeUserModel.Percentages,
			&themeUserModel.TotalLessons,
			&themeUserModel.CompletedLessons,
			&themeUserModel.ThemeID,
			&themeUserModel.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan themesUser: %w", err)
		}

		themeUserModels = append(themeUserModels, themeUserModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	themeUserDomains := themesUserDomainFromModels(themeUserModels)

	return themeUserDomains, nil
}
