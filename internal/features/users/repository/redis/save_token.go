package users_redis_repository

import (
	"context"
	"reflect"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
)

func (r *UserRepository) SaveToken(ctx context.Context, key string, pair *domain.TokenPair) error {
	val := reflect.ValueOf(pair).Elem()

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

		if err := p.Expire(ctx, key, 1*time.Hour); err != nil {
			return err
		}

		return nil
	}

	if _, err := r.pool.Pipelined(ctx, setter); err != nil {
		return err
	}

	return nil
}
