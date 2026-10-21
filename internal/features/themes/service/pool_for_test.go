package themes_service

import (
	"context"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

type postgresRepoStub struct {
	theme domain.Theme
	err   error
	calls int
}

func (r *postgresRepoStub) GetThemes(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Theme, error) {
	return []domain.Theme{}, nil
}

func (r *postgresRepoStub) GetTheme(
	ctx context.Context,
	id int,
) (domain.Theme, error) {
	return r.theme, nil
}

func (r *postgresRepoStub) DeleteTheme(
	ctx context.Context,
	id int,
) error {
	return nil
}

func (r *postgresRepoStub) PatchTheme(
	ctx context.Context,
	id int,
	theme domain.Theme,
) (domain.Theme, error) {
	return domain.Theme{}, nil
}

func (r *postgresRepoStub) CreateTheme(
	ctx context.Context,
	theme domain.Theme,
) (domain.Theme, error) {
	r.calls++
	if r.err != nil {
		return domain.Theme{}, r.err
	}
	return r.theme, nil
}

type redisRepoStub struct {
	err   error
	calls int
	key   string
	theme domain.Theme
}

func (r *redisRepoStub) DeleteToken(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) SaveTheme(
	ctx context.Context,
	key string,
	value domain.Theme,
) error {
	r.calls++
	r.key = key
	r.theme = value
	return r.err
}

func (r *redisRepoStub) GetTheme(
	ctx context.Context,
	key string,
) (domain.Theme, error) {
	return domain.Theme{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) DeleteTheme(
	ctx context.Context,
	key string,
) error {
	return core_errors.ErrNotFound
}

func (r *redisRepoStub) SaveThemes(
	ctx context.Context,
	key string,
	themes []domain.Theme,
) error {
	return nil
}

func (r *redisRepoStub) GetThemes(
	ctx context.Context,
	key string,
) ([]domain.Theme, error) {
	return []domain.Theme{}, core_errors.ErrNotFound
}

func (r *redisRepoStub) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.ThemePatch,
) error {
	return core_errors.ErrNotFound
}
