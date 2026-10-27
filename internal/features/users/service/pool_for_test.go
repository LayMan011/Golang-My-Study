package users_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type postgresRepoStub struct {
	users []domain.User
	err   error
	calls int
}

func (r *postgresRepoStub) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	r.calls++

	if r.err != nil {
		return nil, r.err
	}

	start := 0
	if offset != nil {
		start = *offset
	}

	if start >= len(r.users) {
		return []domain.User{}, nil
	}

	end := len(r.users)
	if limit != nil && start+*limit < end {
		end = start + *limit
	}

	result := make([]domain.User, end-start)
	copy(result, r.users[start:end])

	return result, nil
}

func (r *postgresRepoStub) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, core_errors.ErrNotFound
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
	if r.err != nil {
		return r.err
	}

	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}

	return core_errors.ErrNotFound
}

func (r *postgresRepoStub) PatchUser(
	ctx context.Context,
	id int,
	patch domain.User,
) (domain.User, error) {
	if r.err != nil {
		return domain.User{}, r.err
	}

	for i, user := range r.users {
		if user.ID == id {
			patch.ID = id
			r.users[i] = patch
			return patch, nil
		}
	}

	return domain.User{}, core_errors.ErrNotFound
}

func (r *postgresRepoStub) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	r.calls++
	if r.err != nil {
		return domain.User{}, r.err
	}

	r.users = append(r.users, user)
	return user, nil
}

type redisRepoStub struct {
	err   error
	calls int
	key   string
	user  domain.User

	userCache   map[string]domain.User
	usersCache  map[string][]domain.User
	listVersion int64
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

	if r.err != nil {
		return r.err
	}

	if r.userCache == nil {
		r.userCache = make(map[string]domain.User)
	}
	r.userCache[key] = user

	return nil
}

func (r *redisRepoStub) GetUser(
	ctx context.Context,
	key string,
) (domain.User, error) {
	if r.err != nil {
		return domain.User{}, r.err
	}

	user, ok := r.userCache[key]
	if !ok {
		return domain.User{}, core_errors.ErrNotFound
	}

	return user, nil
}

func (r *redisRepoStub) SaveUsers(
	ctx context.Context,
	key string,
	users []domain.User,
) error {
	if r.err != nil {
		return r.err
	}

	if r.usersCache == nil {
		r.usersCache = make(map[string][]domain.User)
	}

	cached := make([]domain.User, len(users))
	copy(cached, users)
	r.usersCache[key] = cached

	return nil
}

func (r *redisRepoStub) GetUsers(
	ctx context.Context,
	key string,
) ([]domain.User, error) {
	if r.err != nil {
		return nil, r.err
	}

	users, ok := r.usersCache[key]
	if !ok {
		return nil, core_errors.ErrNotFound
	}

	result := make([]domain.User, len(users))
	copy(result, users)

	return result, nil
}

func (r *redisRepoStub) DeleteUser(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) error {
	if r.err != nil {
		return r.err
	}

	for key, user := range r.userCache {
		if user.ID == id {
			delete(r.userCache, key)
			return nil
		}
	}

	return core_errors.ErrNotFound
}

func (r *redisRepoStub) GetUsersListVersion(
	ctx context.Context,
) (int64, error) {
	if r.err != nil {
		return 0, r.err
	}

	return r.listVersion, nil
}

func (r *redisRepoStub) BumpUsersListVersion(
	ctx context.Context,
) error {
	if r.err != nil {
		return r.err
	}

	r.listVersion++
	return nil
}
