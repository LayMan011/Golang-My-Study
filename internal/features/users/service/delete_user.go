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
	s.hash_once[key] = false
	clear(s.hash_many)

	if err := s.userRepositoryRedis.DeleteUser(ctx, key); err != nil {
		return fmt.Errorf("delete user to cache: %w", err)
	}

	key = fmt.Sprintf("token:%d", id)
	s.hash_token[key] = false

	if err := s.userRepositoryRedis.DeleteToken(ctx, key); err != nil {
		return fmt.Errorf("delete token to cache: %w", err)
	}

	return nil
}
