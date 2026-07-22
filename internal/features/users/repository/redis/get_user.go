package users_redis_repository

import (
	"context"
	"fmt"
	"reflect"
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
	v := reflect.ValueOf(&u).Elem()

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

		switch fv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, convErr := strconv.ParseInt(raw, 10, 64)
			if convErr != nil {
				return domain.User{}, fmt.Errorf("field %s: parse int from %q: %w", field.Name, raw, convErr)
			}
			fv.SetInt(n)

		case reflect.String:
			fv.SetString(raw)

		case reflect.Slice:
			if fv.Type().Elem().Kind() == reflect.Uint8 {
				fv.SetBytes([]byte(raw))
			} else {
				return domain.User{}, fmt.Errorf("field %s: unsupported slice type %s", field.Name, fv.Type())
			}

		case reflect.Pointer:
			if fv.Type().Elem().Kind() == reflect.String {
				s := raw
				fv.Set(reflect.ValueOf(&s))
			} else if fv.Type().Elem() == reflect.TypeOf(time.Time{}) {
				// const layout = "2026-02-26T10:30:00Z"
				t, convErr := time.Parse(time.RFC3339Nano, raw)
				if convErr != nil {
					return domain.User{}, fmt.Errorf("field %s: parse time from %q: %w", field.Name, raw, convErr)
				}
				fv.Set(reflect.ValueOf(&t))
			} else {
				return domain.User{}, fmt.Errorf("field %s: unsupported ptr type %s", field.Name, fv.Type())
			}

		case reflect.Struct:
			if fv.Type() == reflect.TypeOf(time.Time{}) {
				// const layout = "2026-02-26T10:30:00Z"
				t, convErr := time.Parse(time.RFC3339Nano, raw)
				if convErr != nil {
					return domain.User{}, fmt.Errorf("field %s: parse time from %q: %w", field.Name, raw, convErr)
				}
				fv.Set(reflect.ValueOf(t))
			} else {
				return domain.User{}, fmt.Errorf("field %s: unsupported struct type %s", field.Name, fv.Type())
			}

		default:
			return domain.User{}, fmt.Errorf("field %s: unsupported kind %s", field.Name, fv.Kind())
		}
	}

	return u, nil
}
