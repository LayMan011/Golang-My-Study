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
	if err := s.userRepositoryPostgres.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	theme_user_ids, err := s.themeUserCachePostgres.GetIDsByThemeID(ctx, id)
	if err != nil {
		return fmt.Errorf("get theme_user IDs for user %d: %w", id, err)
	}

	if err := s.themeUserCacheRedis.InvalidateByIDs(ctx, theme_user_ids); err != nil {
		return fmt.Errorf("invalidate theme_user cache for user %d: %w", id, err)
	}

	if err := s.userRepositoryRedis.BumpUsersListVersion(ctx); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to invalidate users list cache", zap.Error(err))
	}

	if err := s.userRepositoryRedis.DeleteUserSessions(ctx, id); err != nil {
		return fmt.Errorf("revoke user sessions: %w", err)
	}

	key := fmt.Sprintf("user:%d", id)
	if err := s.userRepositoryRedis.DeleteUser(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete user from cache", zap.Error(err))
	}

	return nil
}
