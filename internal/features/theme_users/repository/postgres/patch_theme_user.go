package themes_user_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeUserRepository) PatchThemeUser(
	ctx context.Context,
	id int,
	themeUser domain.ThemeUser,
) (domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE progress.themes_user
	SET 
		completed=$1,
		percentages=$2,
		completed_at=$3,
		version=version + 1
	
	WHERE id=$4 AND version=$5

	RETURNING
		id,
		version,
		completed,
		addition_at,
		completed_at,
		percentages,
		total_lessons,
		completed_lessons,
		theme_id,
		user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		themeUser.Completed,
		themeUser.Percentages,
		themeUser.CompletedAt,
		id,
		themeUser.Version,
	)

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
			return domain.ThemeUser{}, fmt.Errorf(
				"theme with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.ThemeUser{}, fmt.Errorf("scan error: %w", err)
	}

	themeDomain := themeUserDomainFromModel(themeUserModel)

	return themeDomain, nil
}
