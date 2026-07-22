package themes_user_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *ThemeUserRepository) CreateThemeUser(
	ctx context.Context,
	themeUser domain.ThemeUser,
) (domain.ThemeUser, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO progress.themes_user (completed, addition_at, completed_at, percentages, total_lessons, completed_lessons, theme_id, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, version, completed, addition_at, completed_at, percentages, total_lessons, completed_lessons, theme_id, user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		themeUser.Completed,
		themeUser.AdditionAt,
		themeUser.CompletedAt,
		themeUser.Percentages,
		themeUser.TotalLessons,
		themeUser.CompletedLessons,
		themeUser.ThemeID,
		themeUser.UserID,
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
		if errors.Is(err, core_postgres_pool.ErrViolatesForeigKey) {
			return domain.ThemeUser{}, fmt.Errorf(
				"%v: user with id: '%d': %w",
				err,
				themeUser.UserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.ThemeUser{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	themeDomain := themeUserDomainFromModel(themeUserModel)

	return themeDomain, nil
}
