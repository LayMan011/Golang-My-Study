package users_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type postgresRepoStub struct {
	user  domain.User
	err   error
	calls int
}

func (r *postgresRepoStub) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (r *postgresRepoStub) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	return r.user, nil
}

func (r *postgresRepoStub) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	return domain.User{}, nil
}

func (r *postgresRepoStub) DeleteUser(
	ctx context.Context,
	id int,
) error {
	return nil
}

func (r *postgresRepoStub) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	return domain.User{}, nil
}

func (r *postgresRepoStub) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	r.calls++
	if r.err != nil {
		return domain.User{}, r.err
	}
	return r.user, nil
}

type redisRepoStub struct {
	err   error
	calls int
	key   string
	user  domain.User
}

func (r *redisRepoStub) SaveToken(
	ctx context.Context,
	key string,
	pair *domain.TokenPair,
) error {
	return nil
}

func (r *redisRepoStub) DeleteToken(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) GenerateTokenPair(
	ctx context.Context,
	login string,
) (*domain.TokenPair, error) {
	return nil, nil
}

func (r *redisRepoStub) SaveUser(
	ctx context.Context,
	key string,
	user domain.User,
) error {
	r.calls++
	r.key = key
	r.user = user
	return r.err
}

func (r *redisRepoStub) GetUser(
	ctx context.Context,
	key string,
) (domain.User, error) {
	return domain.User{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) DeleteUser(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) SaveUsers(
	ctx context.Context,
	key string,
	users []domain.User,
) error {
	return nil
}

func (r *redisRepoStub) GetUsers(
	ctx context.Context,
	key string,
) ([]domain.User, error) {
	return []domain.User{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) error {
	return core_errors.ErrNotFound
}
