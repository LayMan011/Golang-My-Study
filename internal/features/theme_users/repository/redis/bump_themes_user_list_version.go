package themes_user_redis_repository

import (
	"context"
	"fmt"
)

func (r *ThemeUserRepository) BumpThemesUserListVersion(ctx context.Context) error {
	if _, err := r.pool.Incr(ctx, themesUserListVersionKey); err != nil {
		return fmt.Errorf("bump themes_user list version: %w", err)
	}

	return nil
}
