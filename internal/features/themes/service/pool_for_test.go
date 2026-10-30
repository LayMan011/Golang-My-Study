package themes_service

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
	themes []domain.Theme
	err    error
	calls  int
	state  *stubState
}

func (r *postgresRepoStub) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
	return r.themes, nil
}

func (r *postgresRepoStub) GetTheme(
	ctx context.Context,
	id int,
) (domain.Theme, error) {
	if r.state.unexpectedCacheRead {
		r.state.unexpectedCacheRead = false
		return domain.Theme{}, errors.New("unexpected cache read before Postgres GetTheme")
	}

	for _, theme := range r.themes {
		if theme.ID == id {
			return theme, nil
		}
	}

	return domain.Theme{}, core_errors.ErrNotFound
}

func (r *postgresRepoStub) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	if len(r.themes) > 0 {
		r.themes = r.themes[1:]
		return nil
	}

	return errors.New("failed to delete themes")
}

func (r *postgresRepoStub) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.Theme,
) (domain.Theme, error) {
	for i, theme := range r.themes {
		if theme.ID == id {
			r.themes[i] = patch
			return patch, nil
		}
	}

	return domain.Theme{}, core_errors.ErrNotFound
}

func (r *postgresRepoStub) CreateTheme(
	ctx context.Context,
	theme domain.Theme,
) (domain.Theme, error) {
	r.calls++
	if r.err != nil {
		return domain.Theme{}, r.err
	}

	r.themes = append(r.themes, theme)
	return theme, nil
}

type redisRepoStub struct {
	err   error
	key   string
	theme domain.Theme
	state *stubState

	themeCache  map[string]domain.Theme
	themesCache map[string][]domain.Theme
	listVersion int64
}

func (r *redisRepoStub) SaveTheme(
	ctx context.Context,
	key string,
	value domain.Theme,
) error {
	r.key = key
	r.theme = value

	if r.err != nil {
		return r.err
	}

	if r.themeCache == nil {
		r.themeCache = make(map[string]domain.Theme)
	}

	r.themeCache[key] = value
	return nil
}

func (r *redisRepoStub) GetTheme(
	ctx context.Context,
	key string,
) (domain.Theme, error) {
	if r.err == ErrNotGoToCache {
		r.state.unexpectedCacheRead = true
		return domain.Theme{}, r.err
	}

	if r.err != nil {
		return domain.Theme{}, r.err
	}

	theme, ok := r.themeCache[key]
	if !ok {
		return domain.Theme{}, core_errors.ErrNotFound
	}

	return theme, nil
}

func (r *redisRepoStub) DeleteTheme(
	ctx context.Context,
	key string,
) error {
	if r.err != nil {
		return r.err
	}

	if _, ok := r.themeCache[key]; !ok {
		return core_errors.ErrNotFound
	}

	delete(r.themeCache, key)
	return nil
}

func (r *redisRepoStub) SaveThemes(
	ctx context.Context,
	key string,
	themes []domain.Theme,
) error {
	if r.err != nil {
		return r.err
	}

	if r.themesCache == nil {
		r.themesCache = make(map[string][]domain.Theme)
	}

	cached := make([]domain.Theme, len(themes))
	copy(cached, themes)
	r.themesCache[key] = cached

	return nil
}

func (r *redisRepoStub) GetThemes(
	ctx context.Context,
	key string,
) ([]domain.Theme, error) {
	if r.err != nil {
		return nil, r.err
	}

	themes, ok := r.themesCache[key]
	if !ok {
		return nil, core_errors.ErrNotFound
	}

	result := make([]domain.Theme, len(themes))
	copy(result, themes)

	return result, nil
}

func (r *redisRepoStub) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.ThemePatch,
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
