package themes_redis_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeRepository) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.ThemePatch,
) error {
	data := make(map[string]string)

	if patch.Title.Set && patch.Title.Value != nil {
		data["title"] = *patch.Title.Value
	}

	if patch.Description.Set && patch.Description.Value != nil {
		data["description"] = *patch.Description.Value
	}

	if patch.Subject.Set && patch.Subject.Value != nil {
		data["subject"] = *patch.Subject.Value
	}

	if patch.Level.Set && patch.Level.Value != nil {
		data["level"] = *patch.Level.Value
	}

	if patch.Duration.Set && patch.Duration.Value != nil {
		data["duration"] = *patch.Duration.Value
	}

	if patch.Format.Set && patch.Format.Value != nil {
		data["format"] = *patch.Format.Value
	}

	if len(data) == 0 {
		return nil
	}

	key := fmt.Sprintf("theme:%d", id)
	if err := r.pool.HSet(ctx, key, data); err != nil {
		return fmt.Errorf("patch theme in redis: %w", err)
	}

	return nil
}
