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
	INSERT INTO progress.users (email, created_at, password, full_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, email, created_at, password, full_name;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Email,
		user.CreatedAt,
		user.Password,
		user.FullName,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Email,
		&userModel.CreatedAt,
		&userModel.Password,
		&userModel.FullName,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Email,
		userModel.CreatedAt,
		userModel.Password,
		userModel.FullName,
	)

	return userDomain, nil
}
