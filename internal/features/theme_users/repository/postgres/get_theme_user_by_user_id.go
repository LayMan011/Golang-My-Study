package themes_user_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeUserRepository) GetThemeUserByUserID(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, completed, addition_at, completed_at, percentages, total_lessons, completed_lessons, theme_id, user_id
	FROM progress.themes_user
	WHERE user_id=$1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var themeUserModel ThemeUserModel
	err := row.Scan(
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ThemeUser{}, fmt.Errorf("theme with id='%d': %w", id, core_errors.ErrNotFound)
		}

		return domain.ThemeUser{}, fmt.Errorf("scan error: %w", err)
	}

	themeDomain := themeUserDomainFromModel(themeUserModel)

	return themeDomain, nil
}
