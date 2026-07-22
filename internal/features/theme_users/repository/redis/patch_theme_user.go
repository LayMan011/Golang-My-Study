package themes_user_redis_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeUserRepository) PatchThemeUser(
	ctx context.Context,
	id int,
	patch domain.ThemeUserPatch,
) error {
	data := make(map[string]any)

	if patch.Completed.Set && patch.Completed.Value != nil {
		data["completed"] = *patch.Completed.Value
	}

	if len(data) == 0 {
		return nil
	}

	key := fmt.Sprintf("theme_user:%d", id)

	if err := r.pool.HSet(ctx, key, data); err != nil {
		return fmt.Errorf("patch theme_user in redis: %w", err)
	}

	return nil
}
