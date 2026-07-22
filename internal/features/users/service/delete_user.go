package users_service

import (
	"context"
	"fmt"
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
		return fmt.Errorf("delete user to cache: %w", err)
	}

	key = fmt.Sprintf("token:%d", id)

	if err := s.userRepositoryRedis.DeleteToken(ctx, key); err != nil {
		return fmt.Errorf("delete token to cache: %w", err)
	}

	return nil
}
