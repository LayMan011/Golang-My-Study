package themes_user_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type postgresRepoStub struct {
	themes_user []domain.ThemeUser
	err         error
	calls       int
}

func (r *postgresRepoStub) GetThemesUser(
	ctx context.Context,
	userID *int,
	themeID *int,
	first *int,
	last *int,
) ([]domain.ThemeUser, error) {
	r.calls++

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
	r.calls++
	if r.err != nil {
		return domain.ThemeUser{}, r.err
	}

	r.themes_user = append(r.themes_user, themeUser)
	return themeUser, nil
}

type redisRepoStub struct {
	err        error
	calls      int
	key        string
	theme_user domain.ThemeUser

	themeUserCache  map[string]domain.ThemeUser
	themesUserCache map[string][]domain.ThemeUser
	listVersion     int64
}

func (r *redisRepoStub) SaveThemeUser(
	ctx context.Context,
	key string,
	value domain.ThemeUser,
) error {
	r.calls++
	r.key = key
	r.theme_user = value

	if r.err != nil {
		return r.err
	}

	if r.themeUserCache == nil {
		r.themeUserCache = make(map[string]domain.ThemeUser)
	}
	r.themeUserCache[key] = value

	return nil
}

func (r *redisRepoStub) GetThemeUser(
	ctx context.Context,
	key string,
) (domain.ThemeUser, error) {
	if r.err != nil {
		return domain.ThemeUser{}, r.err
	}

	themeUser, ok := r.themeUserCache[key]
	if !ok {
		return domain.ThemeUser{}, core_errors.ErrNotFound
	}

	return themeUser, nil
}

func (r *redisRepoStub) DeleteThemeUser(
	ctx context.Context,
	key string,
) error {
	if r.err != nil {
		return r.err
	}

	if _, ok := r.themeUserCache[key]; !ok {
		return core_errors.ErrNotFound
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

	if r.themesUserCache == nil {
		r.themesUserCache = make(map[string][]domain.ThemeUser)
	}

	cached := make([]domain.ThemeUser, len(themesUser))
	copy(cached, themesUser)
	r.themesUserCache[key] = cached

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

	for key, themeUser := range r.themeUserCache {
		if themeUser.ID == id {
			delete(r.themeUserCache, key)
			return nil
		}
	}

	return core_errors.ErrNotFound
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
