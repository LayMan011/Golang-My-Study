package users_redis_repository

import (
	"context"
	"fmt"
)

func (r *UserRepository) BumpUsersListVersion(ctx context.Context) error {
	if _, err := r.pool.Incr(ctx, usersListVersionKey); err != nil {
		return fmt.Errorf("bump users list version: %w", err)
	}

	return nil
}
