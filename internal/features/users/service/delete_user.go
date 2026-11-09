package users_service

import (
	"context"
	"fmt"

	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	log := core_logger.FromContext(ctx)

	theme_user_ids, err := s.themeUserCachePostgres.GetIDsByUserID(ctx, id)
	if err != nil {
		log.Warn("failed to get theme_user IDs for cache invalidation", zap.Int("user_id", id), zap.Error(err))
	}

	if err := s.userRepositoryPostgres.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if err := s.themeUserCacheRedis.InvalidateByIDs(ctx, theme_user_ids); err != nil {
		log.Warn("failed to invalidate theme_user cache", zap.Int("theme_id", id), zap.Error(err))
	}

	if err := s.userRepositoryRedis.BumpUsersListVersion(ctx); err != nil {
		log.Warn("failed to invalidate users list cache", zap.Error(err))
	}

	if err := s.userRepositoryRedis.DeleteUserSessions(ctx, id); err != nil {
		return fmt.Errorf("revoke user sessions: %w", err)
	}

	key := fmt.Sprintf("user:%d", id)
	if err := s.userRepositoryRedis.DeleteUser(ctx, key); err != nil {
		log.Warn("failed to delete user from cache", zap.Error(err))
	}

	return nil
}
