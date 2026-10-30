package themes_user_redis_repository

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

const themesUserListVersionKey = "theme_users:list:version"

func (r *ThemeUserRepository) GetThemesUserListVersion(ctx context.Context) (int64, error) {
	value, err := r.pool.Get(ctx, themesUserListVersionKey)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return 0, core_errors.ErrNotFound
		}
		return 0, fmt.Errorf("get themes_user list version: %w", err)
	}

	var version int64
	if _, err := fmt.Sscan(value, &version); err != nil {
		return 0, fmt.Errorf("parse themes_user list version %q: %w", value, err)
	}

	return version, nil
}
