package themes_redis_repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeRepository) DeleteTheme(ctx context.Context, key string) error {
	exists, err := r.pool.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("check theme hash exists: %w", err)
	}
	if exists == 0 {
		return core_errors.ErrNotFound
	}

	var u domain.Theme
	v := reflect.ValueOf(&u).Elem()

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
		return fmt.Errorf("no redis-tagged fields for Theme")
	}

	n, err := r.pool.HDel(ctx, key, fields...).Result()
	if err != nil {
		return fmt.Errorf("delete Theme via HDEL: %w", err)
	}

	if n == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
