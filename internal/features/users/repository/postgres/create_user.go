package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO progress.users (login, password, full_name, phone_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, login, password, full_name, phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Login,
		user.Password,
		user.FullName,
		user.PhoneNumber,
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
