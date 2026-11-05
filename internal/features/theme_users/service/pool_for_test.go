package themes_user_service

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
	themes_user []domain.ThemeUser
	err         error
	state       *stubState

	getThemesUserCalls    int
	createThemesUserCalls int
}

func (r *postgresRepoStub) GetThemesUser(
	ctx context.Context,
	userID *int,
	themeID *int,
	first *int,
	last *int,
) ([]domain.ThemeUser, error) {
	r.getThemesUserCalls++

	if r.err != nil {
		return nil, r.err
	}

	filtered := make([]domain.ThemeUser, 0, len(r.themes_user))
	for _, themeUser := range r.themes_user {
		if userID != nil && themeUser.UserID != *userID {
			continue
		}

		if themeID != nil && themeUser.ThemeID != *themeID {
			continue
		}

		filtered = append(filtered, themeUser)
	}

	start := 0
	if last != nil {
		start = *last
	}

	if start < 0 {
		return nil, core_errors.ErrInvalidArgument
	}

	if start >= len(filtered) {
		return []domain.ThemeUser{}, nil
	}

	end := len(filtered)
	if first != nil {
		if *first < 0 {
			return nil, core_errors.ErrInvalidArgument
		}

		if start+*first < end {
			end = start + *first
		}
	}

	result := make([]domain.ThemeUser, end-start)
	copy(result, filtered[start:end])
	return result, nil
}

func (r *postgresRepoStub) GetThemeUser(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	if r.state.unexpectedCacheRead {
		r.state.unexpectedCacheRead = false
		return domain.ThemeUser{}, errors.New("unexpected cache read before Postgres GetThemeUser")
	}

	for _, theme_user := range r.themes_user {
		if theme_user.ID == id {
			return theme_user, nil
		}
	}

	return domain.ThemeUser{}, core_errors.ErrNotFound
}

func (r *postgresRepoStub) GetThemeUserByUserID(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	return domain.ThemeUser{}, nil
}

func (r *postgresRepoStub) DeleteThemeUser(
	ctx context.Context,
	themeUserID int,
) error {
	if r.err != nil {
		return r.err
	}

	for i, themeUser := range r.themes_user {
		if themeUser.ID == themeUserID {
			r.themes_user = append(r.themes_user[:i], r.themes_user[i+1:]...)
			return nil
		}
	}

	return core_errors.ErrNotFound
}

func (r *postgresRepoStub) PatchThemeUser(
	ctx context.Context,
	id int,
	patch domain.ThemeUser,
) (domain.ThemeUser, error) {
	if r.err != nil {
		return domain.ThemeUser{}, r.err
	}

	for i, themeUser := range r.themes_user {
		if themeUser.ID == id {
			patch.ID = id
			r.themes_user[i] = patch
			return patch, nil
		}
	}

	return domain.ThemeUser{}, core_errors.ErrNotFound
}

func (r *postgresRepoStub) CreateThemeUser(
	ctx context.Context,
	themeUser domain.ThemeUser,
) (domain.ThemeUser, error) {
	r.createThemesUserCalls++
	if r.err != nil {
		return domain.ThemeUser{}, r.err
	}

	r.themes_user = append(r.themes_user, themeUser)
	return themeUser, nil
}

type redisRepoStub struct {
	err   error
	key   string
	state *stubState

	themeUserCache  map[string]domain.ThemeUser
	themesUserCache map[string][]domain.ThemeUser
	listVersion     int64
}

func (r *redisRepoStub) SaveThemeUser(
	ctx context.Context,
	key string,
	value domain.ThemeUser,
) error {
	if r.err != nil {
		return r.err
	}

	version, err := r.GetThemesUserListVersion(ctx)
	if err != nil {
		return err
	}

	r.key = key
	r.themeUserCache[key] = value

	type cacheEntry struct {
		key   string
		users []domain.ThemeUser
	}

	entries := make([]cacheEntry, 0, len(r.themesUserCache))
	for listKey, users := range r.themesUserCache {
		entries = append(entries, cacheEntry{key: listKey, users: users})
	}

	for _, entry := range entries {
		userID, hasUserID := cacheKeyInt(entry.key, "user_id:", "theme_id:")
		themeID, hasThemeID := cacheKeyInt(entry.key, "theme_id:", "limit:")

		if !hasUserID || (userID != -1 && userID != value.UserID) {
			continue
		}

		if !hasThemeID || (themeID != -1 && themeID != value.ThemeID) {
			continue
		}

		limit, hasLimit := cacheKeyInt(entry.key, "limit:", ":offset:")
		offset, hasOffset := cacheKeyInt(entry.key, "offset:", "")
		if !hasLimit || !hasOffset || limit != -1 || offset != -1 {
			continue
		}

		updated := make([]domain.ThemeUser, len(entry.users))
		copy(updated, entry.users)

		found := false
		for i := range updated {
			if updated[i].ID == value.ID {
				updated[i] = value
				found = true
				break
			}
		}

		if !found {
			updated = append(updated, value)
		}

		r.moveThemesUserCacheToVersion(version, entry.key, updated)
	}

	return nil
}

func (r *redisRepoStub) GetThemeUser(
	ctx context.Context,
	key string,
) (domain.ThemeUser, error) {
	if r.err == ErrNotGoToCache {
		r.state.unexpectedCacheRead = true
		return domain.ThemeUser{}, r.err
	}

	if r.err != nil {
		return domain.ThemeUser{}, r.err
	}

	theme_user, ok := r.themeUserCache[key]
	if !ok {
		return domain.ThemeUser{}, core_errors.ErrNotFound
	}

	return theme_user, nil
}

func (r *redisRepoStub) DeleteThemeUser(
	ctx context.Context,
	key string,
) error {
	if r.err != nil {
		return r.err
	}

	themeUser, ok := r.themeUserCache[key]
	if !ok {
		return core_errors.ErrNotFound
	}

	version, err := r.GetThemesUserListVersion(ctx)
	if err != nil {
		return err
	}

	type cacheEntry struct {
		key   string
		users []domain.ThemeUser
	}

	entries := make([]cacheEntry, 0, len(r.themesUserCache))
	for listKey, users := range r.themesUserCache {
		entries = append(entries, cacheEntry{key: listKey, users: users})
	}

	for _, entry := range entries {
		updated := make([]domain.ThemeUser, 0, len(entry.users))
		for _, user := range entry.users {
			if user.ID != themeUser.ID {
				updated = append(updated, user)
			}
		}

		r.moveThemesUserCacheToVersion(version, entry.key, updated)
	}

	delete(r.themeUserCache, key)
	return nil
}

func (r *redisRepoStub) SaveThemesUser(
	ctx context.Context,
	key string,
	themesUser []domain.ThemeUser,
) error {
	if r.err != nil {
		return r.err
	}

	version, err := r.GetThemesUserListVersion(ctx)
	if err != nil {
		return err
	}

	cached := make([]domain.ThemeUser, len(themesUser))
	copy(cached, themesUser)

	r.moveThemesUserCacheToVersion(version, key, cached)
	return nil
}

func (r *redisRepoStub) GetThemesUser(
	ctx context.Context,
	key string,
) ([]domain.ThemeUser, error) {
	if r.err != nil {
		return nil, r.err
	}

	themesUser, ok := r.themesUserCache[key]
	if !ok {
		return nil, core_errors.ErrNotFound
	}

	result := make([]domain.ThemeUser, len(themesUser))
	copy(result, themesUser)

	return result, nil
}

func (r *redisRepoStub) PatchThemeUser(
	ctx context.Context,
	id int,
	patch domain.ThemeUserPatch,
) error {
	if r.err != nil {
		return r.err
	}

	var updatedUser domain.ThemeUser
	found := false

	for _, themeUser := range r.themeUserCache {
		if themeUser.ID != id {
			continue
		}

		if patch.Completed.Set && patch.Completed.Value != nil {
			themeUser.Completed = *patch.Completed.Value
		}

		updatedUser = themeUser
		found = true
		break
	}

	version, err := r.GetThemesUserListVersion(ctx)
	if err != nil {
		return err
	}

	for key, themeUser := range r.themeUserCache {
		if themeUser.ID == id {
			r.themeUserCache[key] = updatedUser
		}
	}

	type cacheEntry struct {
		key   string
		users []domain.ThemeUser
	}

	entries := make([]cacheEntry, 0, len(r.themesUserCache))
	for key, users := range r.themesUserCache {
		entries = append(entries, cacheEntry{key: key, users: users})
	}

	for _, entry := range entries {
		updatedUsers := make([]domain.ThemeUser, len(entry.users))
		copy(updatedUsers, entry.users)

		changed := false
		for i := range updatedUsers {
			if updatedUsers[i].ID == id {
				updatedUsers[i] = updatedUser
				changed = true
			}
		}

		if changed {
			r.moveThemesUserCacheToVersion(version, entry.key, updatedUsers)
		}
	}

	if !found {
		return core_errors.ErrNotFound
	}

	return nil
}

func (r *redisRepoStub) GetThemesUserListVersion(
	ctx context.Context,
) (int64, error) {
	if r.err != nil {
		return 0, r.err
	}

	return r.listVersion, nil
}

func (r *redisRepoStub) BumpThemesUserListVersion(
	ctx context.Context,
) error {
	if r.err != nil {
		return r.err
	}

	r.listVersion++
	return nil
}

func (r *redisRepoStub) moveThemesUserCacheToVersion(
	version int64,
	key string,
	themesUser []domain.ThemeUser,
) {
	const prefix = "themes_user:list:version:"
	const listMarker = "user_id:"

	rest := strings.TrimPrefix(key, prefix)
	if rest == key {
		r.themesUserCache[key] = themesUser
		return
	}

	i := strings.Index(rest, listMarker)
	if i == -1 {
		r.themesUserCache[key] = themesUser
		return
	}

	suffix := rest[i:]
	newKey := prefix + strconv.FormatInt(version, 10) + suffix

	for cachedKey := range r.themesUserCache {
		cachedRest := strings.TrimPrefix(cachedKey, prefix)
		if cachedRest == cachedKey {
			continue
		}

		j := strings.Index(cachedRest, listMarker)
		if j != -1 && cachedRest[j:] == suffix {
			delete(r.themesUserCache, cachedKey)
		}
	}

	r.themesUserCache[newKey] = themesUser
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
