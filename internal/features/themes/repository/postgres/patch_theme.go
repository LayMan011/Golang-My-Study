package themes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeRepository) PatchTheme(
	ctx context.Context,
	id int,
	theme domain.Theme,
) (domain.Theme, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE progress.themes
	SET 
		title=$1,
		description=$2,
		subject=$3,
		level=$4,
		duration=$5,
		format=$6,
		version=version + 1
	
	WHERE id=$7 AND version=$8

	RETURNING
		id,
		version,
		title,
		description,
		created_at,
		subject,
		rating,
		all_ratings,
		number_of_ratings,
		number_of_users,
		price,
		level,
		duration,
		format,
		author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		theme.Title,
		theme.Description,
		theme.Subject,
		theme.Level,
		theme.Duration,
		theme.Format,
		id,
		theme.Version,
	)

	var themeModel ThemeModel
	err := row.Scan(
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
		&themeModel.Level,
		&themeModel.Duration,
		&themeModel.Format,
		&themeModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Theme{}, fmt.Errorf(
				"theme with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Theme{}, fmt.Errorf("scan error: %w", err)
	}

	themeDomain := themeDomainFromModel(themeModel)

	return themeDomain, nil
}
