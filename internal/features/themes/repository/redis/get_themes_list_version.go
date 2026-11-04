package themes_redis_repository

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

const themesListVersionKey = "themes:list:version"

func (r *ThemeRepository) GetThemesListVersion(ctx context.Context) (int64, error) {
	value, err := r.pool.Get(ctx, themesListVersionKey)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			version, incrErr := r.pool.Incr(ctx, themesListVersionKey)
			if incrErr != nil {
				return 0, fmt.Errorf("initialize themes list version: %w", incrErr)
			}

			return version, nil
		}

		return 0, fmt.Errorf("get themes list version: %w", err)
	}

	var version int64
	if _, err := fmt.Sscan(value, &version); err != nil {
		return 0, fmt.Errorf("parse themes list version %q: %w", value, err)
	}

	return version, nil
}
