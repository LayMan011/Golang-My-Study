package themes_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeRepository) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM progress.themes
	WHERE id=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("theme with id='%d': %w", id, core_errors.ErrNotFound)
	}

	return nil
}
