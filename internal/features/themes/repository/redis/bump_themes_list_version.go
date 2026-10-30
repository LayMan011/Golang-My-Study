package themes_redis_repository

import (
	"context"
	"fmt"
)

func (r *ThemeRepository) BumpThemesListVersion(ctx context.Context) error {
	if _, err := r.pool.Incr(ctx, themesListVersionKey); err != nil {
		return fmt.Errorf("bump themes list version: %w", err)
	}

	return nil
}
