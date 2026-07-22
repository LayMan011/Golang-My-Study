package users_redis_repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *UserRepository) GetToken(ctx context.Context, key string) (domain.TokenPair, error) {
	data, err := r.pool.HGetAll(ctx, key)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("get token from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.TokenPair{}, core_errors.ErrNotFound
	}

	var t domain.TokenPair
	v := reflect.ValueOf(&t).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("redis")
		if tag == "" || tag == "-" {
			continue
		}

		raw, ok := data[tag]
		if !ok {
			continue
		}

		fv := v.Field(i)
		if !fv.CanSet() {
			continue
		}

		fv.SetString(raw)
	}

	return t, nil
}
