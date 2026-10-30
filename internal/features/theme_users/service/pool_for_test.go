package themes_user_service

import (
	"context"
	"errors"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type stubState struct {
	unexpectedCacheRead bool
}

type postgresRepoStub struct {
	themes_user []domain.ThemeUser
	err         error
	calls       int
	state       *stubState
}

func (r *postgresRepoStub) GetThemesUser(
	ctx context.Context,
	userID *int,
	themeID *int,
	first *int,
	last *int,
) ([]domain.ThemeUser, error) {
	return r.themes_user, nil
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
	if len(r.themes_user) > 0 {
		r.themes_user = r.themes_user[1:]
		return nil
	}

	return errors.New("failed to delete themes_user")
}

func (r *postgresRepoStub) PatchThemeUser(
	ctx context.Context,
	id int,
	patch domain.ThemeUser,
) (domain.ThemeUser, error) {
	for i, theme_user := range r.themes_user {
		if theme_user.ID == id {
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
	state      *stubState

	themeUserCache  map[string]domain.ThemeUser
	themesUserCache map[string][]domain.ThemeUser
	listVersion     int64
}

func (r *redisRepoStub) DeleteToken(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) SaveThemeUser(
	ctx context.Context,
	key string,
	value domain.ThemeUser,
) error {
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
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) GetThemesListVersion(
	ctx context.Context,
) (int64, error) {
	if r.err != nil {
		return 0, r.err
	}

	return r.listVersion, nil
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
