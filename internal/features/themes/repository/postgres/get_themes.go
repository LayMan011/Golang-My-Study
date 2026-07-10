package themes_postgres_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeRepository) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, created_at, subject, rating, all_ratings, number_of_ratings, number_of_users, price, author_user_id
	FROM progress.themes
	%s
	ORDER by id ASC
	LIMIT $1
	OFFSET $2;
	`
	args := []any{limit, offset}

	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id=$3")
		args = append(args, userID)
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

	var themeModels []ThemeModel

	for rows.Next() {
		var themeModel ThemeModel

		err := rows.Scan(
			&themeModel.ID,
			&themeModel.Version,
			&themeModel.Title,
			&themeModel.Description,
			&themeModel.CreatedAt,
			&themeModel.Subject,
			&themeModel.Rating,
			&themeModel.AllRatings,
			&themeModel.NumberOfRatings,
			&themeModel.NumberOfUsers,
			&themeModel.Price,
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
