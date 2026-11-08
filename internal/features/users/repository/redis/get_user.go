package users_redis_repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) GetUser(ctx context.Context, key string) (domain.User, error) {
	data, err := r.pool.HGetAll(ctx, key)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.User{}, core_errors.ErrNotFound
	}

	var u domain.User

	if raw, ok := data["id"]; ok {
		u.ID, err = strconv.Atoi(raw)
		if err != nil {
			return domain.User{}, fmt.Errorf("field ID: parse int from %q: %w", raw, err)
		}
	}

	if raw, ok := data["version"]; ok {
		u.Version, err = strconv.Atoi(raw)
		if err != nil {
			return domain.User{}, fmt.Errorf("field Version: parse int from %q: %w", raw, err)
		}
	}

	if raw, ok := data["email"]; ok {
		u.Email = raw
	}

	if raw, ok := data["created_at"]; ok {
		u.CreatedAt, err = time.Parse(time.RFC3339Nano, raw)
		if err != nil {
			return domain.User{}, fmt.Errorf("field CreatedAt: parse time from %q: %w", raw, err)
		}
	}

	if raw, ok := data["password"]; ok {
		u.Password = []byte(raw)
	}

	if raw, ok := data["full_name"]; ok {
		u.FullName = raw
	}

	return u, nil
}
