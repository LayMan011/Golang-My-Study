package users_redis_repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) DeleteToken(ctx context.Context, key string) error {
	exists, err := r.pool.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("check token hash exists: %w", err)
	}
	if exists == 0 {
		return core_errors.ErrNotFound
	}

	var t domain.TokenPair
	v := reflect.ValueOf(&t).Elem()

	fields := make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("redis")
		if tag == "" || tag == "-" {
			continue
		}
		fields = append(fields, tag)
	}

	if len(fields) == 0 {
		return fmt.Errorf("no redis-tagged fields for TokenPair")
	}

	deleted, err := r.pool.HDel(ctx, key, fields...).Result()
	if err != nil {
		return fmt.Errorf("delete token via HDEL: %w", err)
	}

	if deleted == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
