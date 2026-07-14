package users_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type UserService struct {
	userRepositoryPostgres UserRepositoryPostgres
	userRepositoryRedis    UserRepositoryRedis
}

type UserRepositoryRedis interface {
	SaveToken(
		ctx context.Context,
		login string,
		pair *domain.TokenPair,
	) error
}

type UserRepositoryPostgres interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	GetUserByLogin(
		ctx context.Context,
		login string,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	userRepository UserRepositoryPostgres,
) *UserService {
	return &UserService{
		userRepositoryPostgres: userRepository,
	}
}
