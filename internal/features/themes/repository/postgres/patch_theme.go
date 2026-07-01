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

	fmt.Println(theme)

	query := `
	UPDATE progress.themes
	SET 
		title=$1,
		description=$2,
		completed=$3,
		completed_at=$4,
		percentages=$5,
		version=version + 1
	
	WHERE id=$6 AND version=$7

	RETURNING
		id,
		version,
		title,
		description,
		completed,
		created_at,
		completed_at,
		percentages,
		author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		theme.Title,
		theme.Description,
		theme.Completed,
		theme.CompletedAt,
		theme.Percentages,
		id,
		theme.Version,
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
