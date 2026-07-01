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
	INSERT INTO progress.themes (title, description, completed, created_at, completed_at, percentages, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, version, title, description, completed, created_at, completed_at, percentages, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		theme.Title,
		theme.Description,
		theme.Completed,
		theme.CreatedAt,
		theme.CompletedAt,
		theme.Percentages,
		theme.AuthorUserID,
	)

	var themeModel ThemeModel
	err := row.Scan(
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
