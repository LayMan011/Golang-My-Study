package themes_user_redis_repository

import (
	"context"
	"reflect"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *ThemeUserRepository) SaveThemeUser(ctx context.Context, key string, value domain.ThemeUser) error {
	val := reflect.ValueOf(value)

	setter := func(p core_redis_pool.Pipeline) error {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			tag := field.Tag.Get("redis")
			if tag == "" {
				continue
			}

			if err := p.HSet(ctx, key, tag, val.Field(i).Interface()); err != nil {
				return err
			}
		}

		if err := p.Expire(ctx, key, core_redis_pool.TTLDataSave); err != nil {
			return err
		}

		return nil
	}

	if _, err := r.pool.Pipelined(ctx, setter); err != nil {
		return err
	}

	return nil
}
