package users_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

type UserService struct {
	userRepositoryPostgres UserRepositoryPostgres
	userRepositoryRedis    UserRepositoryRedis
	themeUserCachePostgres ThemeUserCachePostgres
	themeUserCacheRedis    ThemeUserCacheRedis
}

type ThemeUserCachePostgres interface {
	GetIDsByUserID(
		ctx context.Context,
		userID int,
	) ([]int, error)
}

type ThemeUserCacheRedis interface {
	InvalidateByIDs(
		ctx context.Context,
		ids []int,
	) error
}

type UserRepositoryRedis interface {
	SaveToken(
		ctx context.Context,
		key string,
		pair *domain.TokenPair,
	) error

	DeleteToken(
		ctx context.Context,
		key string,
	) error

	GenerateTokenPair(
		ctx context.Context,
		login string,
		userID int,
	) (*domain.TokenPair, error)

	DeleteUserSessions(
		ctx context.Context,
		userID int,
	) error

	SaveUser(
		ctx context.Context,
		key string,
		value domain.User,
	) error

	GetUser(
		ctx context.Context,
		key string,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		key string,
	) error

	SaveUsers(
		ctx context.Context,
		key string,
		users []domain.User,
	) error

	GetUsers(
		ctx context.Context,
		key string,
	) ([]domain.User, error)

	PatchUser(
		ctx context.Context,
		id int,
		patch domain.UserPatch,
	) error

	GetUsersListVersion(
		ctx context.Context,
	) (int64, error)

	BumpUsersListVersion(
		ctx context.Context,
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

	GetUserByEmail(
		ctx context.Context,
		email string,
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
	userRepositoryPostgres UserRepositoryPostgres,
	userRepositoryRedis UserRepositoryRedis,
	themeUserCachePostgres ThemeUserCachePostgres,
	themeUserCacheRedis ThemeUserCacheRedis,
) *UserService {
	return &UserService{
		userRepositoryPostgres: userRepositoryPostgres,
		userRepositoryRedis:    userRepositoryRedis,
		themeUserCachePostgres: themeUserCachePostgres,
		themeUserCacheRedis:    themeUserCacheRedis,
	}
}
