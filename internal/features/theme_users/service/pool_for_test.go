package themes_user_service

import (
	"context"
	"errors"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type postgresRepoStub struct {
	theme_user  domain.ThemeUser
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
	return r.themes_user, nil
}

func (r *postgresRepoStub) GetThemeUser(
	ctx context.Context,
	id int,
) (domain.ThemeUser, error) {
	return r.theme_user, nil
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
	return r.theme_user, nil
}

type redisRepoStub struct {
	err        error
	calls      int
	key        string
	theme_user domain.ThemeUser
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
	r.calls++
	r.key = key
	r.theme_user = value
	return r.err
}

func (r *redisRepoStub) GetThemeUser(
	ctx context.Context,
	key string,
) (domain.ThemeUser, error) {
	return domain.ThemeUser{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) DeleteThemeUser(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) SaveThemesUser(
	ctx context.Context,
	key string,
	themesUser []domain.ThemeUser,
) error {
	return nil
}

func (r *redisRepoStub) GetThemesUser(
	ctx context.Context,
	key string,
) ([]domain.ThemeUser, error) {
	return []domain.ThemeUser{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) PatchThemeUser(
	ctx context.Context,
	id int,
	patch domain.ThemeUserPatch,
) error {
	return core_errors.ErrNotFound
}
