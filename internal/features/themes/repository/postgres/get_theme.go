package themes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeRepository) GetTheme(
	ctx context.Context,
	id int,
) (domain.Theme, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, created_at, subject, rating, all_ratings, number_of_ratings, number_of_users, price, author_user_id
	FROM progress.themes
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

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
		&themeModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Theme{}, fmt.Errorf("theme with id='%d': %w", id, core_errors.ErrNotFound)
		}

		return domain.Theme{}, fmt.Errorf("scan error: %w", err)
	}

	themeDomain := themeDomainFromModel(themeModel)

	return themeDomain, nil
}
