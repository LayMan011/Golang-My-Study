package themes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeRepository) CreateTheme(
	ctx context.Context,
	theme domain.Theme,
) (domain.Theme, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO progress.themes (title, description, created_at, subject, rating, all_ratings, number_of_ratings, number_of_users, price, level, duration, format, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id, version, title, description, created_at, subject, rating, all_ratings, number_of_ratings, number_of_users, price, level, duration, format, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		theme.Title,
		theme.Description,
		theme.CreatedAt,
		theme.Subject,
		theme.Rating,
		theme.AllRatings,
		theme.NumberOfRatings,
		theme.NumberOfUsers,
		theme.Price,
		theme.Level,
		theme.Duration,
		theme.Format,
		theme.AuthorUserID,
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
		if errors.Is(err, core_postgres_pool.ErrViolatesForeigKey) {
			return domain.Theme{}, fmt.Errorf(
				"%v: user with id: '%d': %w",
				err,
				theme.AuthorUserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Theme{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	themeDomain := themeDomainFromModel(themeModel)

	return themeDomain, nil
}
