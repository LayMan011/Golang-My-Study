package users_service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type stubState struct {
	unexpectedCacheRead bool
}

type postgresRepoStub struct {
	users []domain.User
	err   error
	state *stubState

	getUsersCalls   int
	createUserCalls int
}

func (r *postgresRepoStub) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	r.getUsersCalls++

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
	if r.state.unexpectedCacheRead {
		r.state.unexpectedCacheRead = false
		return domain.User{}, errors.New("unexpected cache read before Postgres GetUser")
	}

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
	r.createUserCalls++
	if r.err != nil {
		return domain.User{}, r.err
	}

	r.users = append(r.users, user)
	return user, nil
}

type redisRepoStub struct {
	err   error
	key   string
	state *stubState

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
	if r.err != nil {
		return r.err
	}

	version, err := r.GetUsersListVersion(ctx)
	if err != nil {
		return err
	}

	r.key = key
	r.userCache[key] = user

	type cacheEntry struct {
		key   string
		users []domain.User
	}

	entries := make([]cacheEntry, 0, len(r.usersCache))
	for listKey, users := range r.usersCache {
		entries = append(entries, cacheEntry{key: listKey, users: users})
	}

	for _, entry := range entries {
		if !isUnpaginatedUsersList(entry.key) {
			continue
		}

		updated := make([]domain.User, len(entry.users))
		copy(updated, entry.users)

		found := false
		for i := range updated {
			if updated[i].ID == user.ID {
				updated[i] = user
				found = true
				break
			}
		}

		if !found {
			updated = append(updated, user)
		}

		r.moveUsersCacheToVersion(version, entry.key, updated)
	}

	return nil
}

func (r *redisRepoStub) GetUser(
	ctx context.Context,
	key string,
) (domain.User, error) {
	if r.err == ErrNotGoToCache {
		r.state.unexpectedCacheRead = true
		return domain.User{}, r.err
	}

	user, ok := r.userCache[key]
	if !ok {
		return domain.User{}, core_errors.ErrNotFound
	}

	return user, nil
}

func (r *redisRepoStub) DeleteUser(
	ctx context.Context,
	key string,
) error {
	if r.err != nil {
		return r.err
	}

	user, ok := r.userCache[key]
	if !ok {
		return core_errors.ErrNotFound
	}

	version, err := r.GetUsersListVersion(ctx)
	if err != nil {
		return err
	}

	type cacheEntry struct {
		key   string
		users []domain.User
	}

	entries := make([]cacheEntry, 0, len(r.usersCache))
	for listKey, users := range r.usersCache {
		entries = append(entries, cacheEntry{key: listKey, users: users})
	}

	delete(r.userCache, key)

	for _, entry := range entries {
		if !isUnpaginatedUsersList(entry.key) {
			delete(r.usersCache, entry.key)
			continue
		}

		updated := make([]domain.User, 0, len(entry.users))
		for _, cachedUser := range entry.users {
			if cachedUser.ID != user.ID {
				updated = append(updated, cachedUser)
			}
		}

		r.moveUsersCacheToVersion(version, entry.key, updated)
	}

	return nil
}

func (r *redisRepoStub) SaveUsers(
	ctx context.Context,
	key string,
	users []domain.User,
) error {
	if r.err != nil {
		return r.err
	}

	version, err := r.GetUsersListVersion(ctx)
	if err != nil {
		return err
	}

	cached := make([]domain.User, len(users))
	copy(cached, users)

	r.moveUsersCacheToVersion(version, key, cached)
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

func (r *redisRepoStub) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) error {
	if r.err != nil {
		return r.err
	}

	found := false

	for key, user := range r.userCache {
		if user.ID != id {
			continue
		}

		user.ApplyPatch(patch)
		r.userCache[key] = user
		found = true
	}

	if !found {
		return core_errors.ErrNotFound
	}

	return nil
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

const usersListCachePrefix = "users:list:version:"

func (r *redisRepoStub) moveUsersCacheToVersion(
	version int64,
	key string,
	users []domain.User,
) {
	rest := strings.TrimPrefix(key, usersListCachePrefix)
	if rest == key {
		r.usersCache[key] = users
		return
	}

	i := 0
	for i < len(rest) && rest[i] >= '0' && rest[i] <= '9' {
		i++
	}
	if i == 0 {
		r.usersCache[key] = users
		return
	}

	suffix := rest[i:]
	newKey := usersListCachePrefix + strconv.FormatInt(version, 10) + suffix

	for cachedKey := range r.usersCache {
		cachedRest := strings.TrimPrefix(cachedKey, usersListCachePrefix)
		if cachedRest == cachedKey {
			continue
		}

		j := 0
		for j < len(cachedRest) && cachedRest[j] >= '0' && cachedRest[j] <= '9' {
			j++
		}

		if j > 0 && cachedRest[j:] == suffix {
			delete(r.usersCache, cachedKey)
		}
	}

	r.usersCache[newKey] = users
}

func isUnpaginatedUsersList(key string) bool {
	limit, hasLimit := cacheKeyInt(key, "limit:", ":offset:")
	offset, hasOffset := cacheKeyInt(key, "offset:", "")

	if !hasLimit && !hasOffset {
		return true
	}

	return hasLimit && hasOffset && limit == -1 && offset == -1
}

func cacheKeyInt(key, marker, nextMarker string) (int, bool) {
	start := strings.Index(key, marker)
	if start == -1 {
		return 0, false
	}
	start += len(marker)

	end := len(key)
	if nextMarker != "" {
		if i := strings.Index(key[start:], nextMarker); i >= 0 {
			end = start + i
		}
	}

	value, err := strconv.Atoi(key[start:end])
	if err != nil {
		return 0, false
	}
	return value, true
}

func (r *redisRepoStub) BumpThemesListVersion(
	ctx context.Context,
) error {
	if r.err != nil {
		return r.err
	}

	r.listVersion++
	return nil
}
