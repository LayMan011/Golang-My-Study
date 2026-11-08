package users_redis_repository

import (
	"context"
	"fmt"
)

func (r *UserRepository) DeleteUserSessions(
	ctx context.Context,
	userID int,
) error {
	indexKey := fmt.Sprintf("user_sessions:%d", userID)

	sessions, err := r.pool.HGetAll(ctx, indexKey)
	if err != nil {
		return fmt.Errorf("get user sessions: %w", err)
	}

	keys := make([]string, 0, len(sessions)+1)
	for sessionKey := range sessions {
		keys = append(keys, sessionKey)
	}
	keys = append(keys, indexKey)

	if err := r.pool.Del(ctx, keys...); err != nil {
		return fmt.Errorf("delete user sessions: %w", err)
	}

	return nil
}
