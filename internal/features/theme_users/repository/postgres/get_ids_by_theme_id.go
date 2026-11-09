package themes_user_postgres_repository

import (
	"context"
	"fmt"
)

func (r *ThemeUserRepository) GetIDsByThemeID(
	ctx context.Context,
	themeID int,
) ([]int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id
	FROM progress.themes_user
	WHERE theme_id = $1
	ORDER by id ASC;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		themeID,
	)
	if err != nil {
		return nil, fmt.Errorf("select themes_user: %w", err)
	}
	defer rows.Close()

	ids := make([]int, 0)
	for rows.Next() {
		var id int

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan ids: %w", err)
		}

		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return ids, nil
}
