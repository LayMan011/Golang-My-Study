package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *UserRepository) GetUserByLogin(
	ctx context.Context,
	login string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, login, password, full_name, phone_number
	FROM progress.users
	WHERE login=$1;
	`

	row := r.pool.QueryRow(ctx, query, login)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Login,
		&userModel.Password,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with login='%s': %w",
				login,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Login,
		userModel.Password,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
