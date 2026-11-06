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
	limit *int,
	offset *int,
) ([]domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, completed, addition_at, completed_at, percentages, total_lessons, completed_lessons, theme_id, user_id
	FROM progress.themes_user
	%s
	ORDER by id ASC
	%s;
	`

	add_strings := [2]string{}
	args := []any{}

	if userID != nil && themeID != nil {
		add_strings[0] = `WHERE user_id=$1 AND theme_id=$2`
		args = append(args, userID)
		args = append(args, themeID)
	} else if userID != nil && themeID == nil {
		add_strings[0] = `WHERE user_id=$1`
		args = append(args, userID)
	} else if userID == nil && themeID != nil {
		add_strings[0] = `WHERE theme_id=$1`
		args = append(args, themeID)
	} else {
		add_strings[0] = ``
	}

	if limit != nil && offset != nil {
		args = append(args, limit, offset)
		add_strings[1] = fmt.Sprintf(`LIMIT $%d OFFSET $%d`, len(args)-1, len(args))
	} else if limit != nil && offset == nil {
		args = append(args, limit)
		add_strings[1] = fmt.Sprintf(`LIMIT $%d`, len(args))
	} else if limit == nil && offset != nil {
		args = append(args, offset)
		add_strings[1] = fmt.Sprintf(`OFFSET $%d`, len(args))
	} else {
		add_strings[1] = ``
	}

	query = fmt.Sprintf(query, add_strings[0], add_strings[1])

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select themes_user: %w", err)
	}
	defer rows.Close()

	themeUserModels := make([]ThemeUserModel, 0)
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
