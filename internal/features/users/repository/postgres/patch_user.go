package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"
)

func (r *UserRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE progress.users
	SET
		password=$1,
		full_name=$2,
		phone_number=$3,
		version=version+1
	WHERE id=$4 AND version=$5
	RETURNING
		id,
		version,
		login,
		password,
		full_name,
		phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Password,
		user.FullName,
		user.PhoneNumber,
		id,
		user.Version,
	)

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
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
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
