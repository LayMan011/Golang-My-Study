package users_service

import (
	"context"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (s *UserService) GetUserByLogin(
	ctx context.Context,
	login string,
) (domain.User, error) {
	user, err := s.userRepositoryPostgres.GetUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
