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

	key := fmt.Sprintf("user:%d", id)

	if err := s.userRepositoryRedis.DeleteUser(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete user from cache", zap.Error(err))
	}

	key = fmt.Sprintf("token:%d", id)

	if err := s.userRepositoryRedis.DeleteToken(ctx, key); err != nil {
		log := core_logger.FromContext(ctx)
		log.Warn("failed to delete token from cache", zap.Error(err))
	}

	return nil
}
