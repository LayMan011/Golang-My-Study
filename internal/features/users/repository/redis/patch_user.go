package users_redis_repository

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) error {
	data := make(map[string]any)

	if patch.Password.Set && patch.Password.Value != nil {
		data["password"] = string(*patch.Password.Value)
	}

	if patch.FullName.Set && patch.FullName.Value != nil {
		data["full_name"] = patch.FullName.Value
	}

	if len(data) == 0 {
		return nil
	}

	key := fmt.Sprintf("user:%d", id)

	if err := r.pool.HSet(ctx, key, data); err != nil {
		return fmt.Errorf("patch user in redis: %w", err)
	}

	return nil
}
