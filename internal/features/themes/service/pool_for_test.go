package themes_service

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
	themes []domain.Theme
	err    error
	state  *stubState

	getThemesCalls   int
	createThemeCalls int
}

func (r *postgresRepoStub) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
	r.getThemesCalls++

	if r.err != nil {
		return nil, r.err
	}

	filtered := make([]domain.Theme, 0, len(r.themes))
	for _, theme := range r.themes {
		if userID != nil && theme.AuthorUserID != *userID {
			continue
		}

		filtered = append(filtered, theme)
	}

	start := 0
	if offset != nil {
		start = *offset
	}

	if start >= len(filtered) {
		return []domain.Theme{}, nil
	}

	end := len(filtered)
	if limit != nil && start+*limit < end {
		end = start + *limit
	}

	result := make([]domain.Theme, end-start)
	copy(result, filtered[start:end])

	return result, nil
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
	if r.err != nil {
		return r.err
	}

	for i, theme := range r.themes {
		if theme.ID == id {
			r.themes = append(r.themes[:i], r.themes[i+1:]...)
			return nil
		}
	}

	return core_errors.ErrNotFound
}

func (r *postgresRepoStub) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.Theme,
) (domain.Theme, error) {
	if r.err != nil {
		return domain.Theme{}, r.err
	}

	for i, theme := range r.themes {
		if theme.ID == id {
			patch.ID = id
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
	r.createThemeCalls++
	if r.err != nil {
		return domain.Theme{}, r.err
	}

	r.themes = append(r.themes, theme)
	return theme, nil
}

type redisRepoStub struct {
	err   error
	key   string
	state *stubState

	themeCache  map[string]domain.Theme
	themesCache map[string][]domain.Theme
	listVersion int64
}

const themesListCachePrefix = "themes:list:version:"

func (r *redisRepoStub) SaveTheme(
	ctx context.Context,
	key string,
	value domain.Theme,
) error {
	if r.err != nil {
		return r.err
	}

	version, err := r.GetThemesListVersion(ctx)
	if err != nil {
		return err
	}

	r.key = key
	r.themeCache[key] = value

	type cacheEntry struct {
		key    string
		themes []domain.Theme
	}

	entries := make([]cacheEntry, 0, len(r.themesCache))
	for listKey, themes := range r.themesCache {
		entries = append(entries, cacheEntry{key: listKey, themes: themes})
	}

	for _, entry := range entries {
		if !isUnpaginatedThemesList(entry.key) {
			continue
		}

		updated := make([]domain.Theme, len(entry.themes))
		copy(updated, entry.themes)

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

		r.moveThemesCacheToVersion(version, entry.key, updated)
	}

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

	theme, ok := r.themeCache[key]
	if !ok {
		return core_errors.ErrNotFound
	}

	version, err := r.GetThemesListVersion(ctx)
	if err != nil {
		return err
	}

	type cacheEntry struct {
		key    string
		themes []domain.Theme
	}

	entries := make([]cacheEntry, 0, len(r.themesCache))
	for listKey, themes := range r.themesCache {
		entries = append(entries, cacheEntry{key: listKey, themes: themes})
	}

	delete(r.themeCache, key)

	for _, entry := range entries {
		updated := make([]domain.Theme, 0, len(entry.themes))
		for _, cachedTheme := range entry.themes {
			if cachedTheme.ID != theme.ID {
				updated = append(updated, cachedTheme)
			}
		}

		r.moveThemesCacheToVersion(version, entry.key, updated)
	}

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

	version, err := r.GetThemesListVersion(ctx)
	if err != nil {
		return err
	}

	if r.themesCache == nil {
		r.themesCache = make(map[string][]domain.Theme)
	}

	cached := make([]domain.Theme, len(themes))
	copy(cached, themes)

	r.moveThemesCacheToVersion(version, key, cached)
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
	if r.err != nil {
		return r.err
	}

	found := false

	for key, theme := range r.themeCache {
		if theme.ID != id {
			continue
		}

		if err := theme.ApplyPatch(patch); err != nil {
			return err
		}

		r.themeCache[key] = theme
		found = true
	}

	type cacheEntry struct {
		key    string
		themes []domain.Theme
	}

	version, err := r.GetThemesListVersion(ctx)
	if err != nil {
		return err
	}

	entries := make([]cacheEntry, 0, len(r.themesCache))
	for key, themes := range r.themesCache {
		entries = append(entries, cacheEntry{key: key, themes: themes})
	}

	for _, entry := range entries {
		themes := make([]domain.Theme, len(entry.themes))
		copy(themes, entry.themes)

		changed := false
		for i := range themes {
			if themes[i].ID != id {
				continue
			}

			if err := themes[i].ApplyPatch(patch); err != nil {
				return err
			}

			found = true
			changed = true
		}

		if changed || entry.key != "" {
			r.moveThemesCacheToVersion(version, entry.key, themes)
		}
	}

	if !found {
		return core_errors.ErrNotFound
	}

	return nil
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

func (r *redisRepoStub) moveThemesCacheToVersion(
	version int64,
	key string,
	themes []domain.Theme,
) {
	rest := strings.TrimPrefix(key, themesListCachePrefix)
	if rest == key {
		r.themesCache[key] = themes
		return
	}

	i := 0
	for i < len(rest) && rest[i] >= '0' && rest[i] <= '9' {
		i++
	}
	if i == 0 {
		r.themesCache[key] = themes
		return
	}

	suffix := rest[i:]
	newKey := themesListCachePrefix + strconv.FormatInt(version, 10) + suffix

	for cachedKey := range r.themesCache {
		cachedRest := strings.TrimPrefix(cachedKey, themesListCachePrefix)
		if cachedRest == cachedKey {
			continue
		}

		j := 0
		for j < len(cachedRest) && cachedRest[j] >= '0' && cachedRest[j] <= '9' {
			j++
		}

		if j > 0 && cachedRest[j:] == suffix {
			delete(r.themesCache, cachedKey)
		}
	}

	r.themesCache[newKey] = themes
}

func isUnpaginatedThemesList(key string) bool {
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
