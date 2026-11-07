package users_redis_repository

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

const usersListVersionKey = "users:list:version"

func (r *UserRepository) GetUsersListVersion(ctx context.Context) (int64, error) {
	value, err := r.pool.Get(ctx, usersListVersionKey)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			version, incrErr := r.pool.Incr(ctx, usersListVersionKey)
			if incrErr != nil {
				return 0, fmt.Errorf("initialize users list version: %w", incrErr)
			}

			return version, nil
		}

		return 0, fmt.Errorf("get users list version: %w", err)
	}

	var version int64
	if _, err := fmt.Sscan(value, &version); err != nil {
		return 0, fmt.Errorf("parse users list version %q: %w", value, err)
	}

	return version, nil
}
