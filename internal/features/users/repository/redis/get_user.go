package users_redis_repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) GetUserFromHash(ctx context.Context, login string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	if err := r.pool.DoError(ctx, "SELECT", 0); err != nil {
		return nil, fmt.Errorf("db switch error: %w", err)
	}

	key := fmt.Sprintf("user:%s", login)

	data, err := r.pool.HGetAllResult(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("user with id %s not found", login)
	}

	userID, err := strconv.Atoi(data["id"])
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	version, err := strconv.Atoi(data["version"])
	if err != nil {
		return nil, fmt.Errorf("failed to parse version: %w", err)
	}

	password := []byte(data["password"])

	var phoneNumber *string
	if phoneVal, ok := data["phone_number"]; ok && phoneVal != "" {
		phoneNumber = &phoneVal
	}

	user := domain.NewUser(userID, version, data["login"], password, data["fullname"], phoneNumber)

	return &user, nil
}
