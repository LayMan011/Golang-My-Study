package themes_user_redis_repository

import (
	"context"
	"fmt"
)

func (r *ThemeUserRepository) InvalidateByIDs(
	ctx context.Context,
	ids []int,
) error {
	if len(ids) > 0 {
		keys := make([]string, 0, len(ids))
		for _, id := range ids {
			keys = append(keys, fmt.Sprintf("theme_user:%d", id))
		}

		if err := r.pool.Del(ctx, keys...); err != nil {
			return fmt.Errorf("delete theme_user cache keys: %w", err)
		}
	}

	if err := r.BumpThemesUserListVersion(ctx); err != nil {
		return fmt.Errorf("bump themes_user list version: %w", err)
	}

	return nil
}
