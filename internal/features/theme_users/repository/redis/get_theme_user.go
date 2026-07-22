package themes_user_redis_repository

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeUserRepository) GetThemeUser(ctx context.Context, key string) (domain.ThemeUser, error) {
	data, err := r.pool.HGetAll(ctx, key)
	if err != nil {
		return domain.ThemeUser{}, fmt.Errorf("get theme_user from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.ThemeUser{}, core_errors.ErrNotFound
	}

	var tu domain.ThemeUser
	v := reflect.ValueOf(&tu).Elem()

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
				return domain.ThemeUser{}, fmt.Errorf("field %s: parse int from %q: %w", field.Name, raw, convErr)
			}
			fv.SetInt(n)

		case reflect.Bool:
			b, convErr := strconv.ParseBool(raw)
			if convErr != nil {
				return domain.ThemeUser{}, fmt.Errorf("field %s: parse bool from %q: %w", field.Name, raw, convErr)
			}
			fv.SetBool(b)

		case reflect.String:
			fv.SetString(raw)

		case reflect.Struct:
			if fv.Type() == reflect.TypeOf(time.Time{}) {
				const layout = "2026-02-26T10:30:00Z"
				tm, convErr := time.Parse(layout, raw)
				if convErr != nil {
					return domain.ThemeUser{}, fmt.Errorf("field %s: parse time from %q: %w", field.Name, raw, convErr)
				}
				fv.Set(reflect.ValueOf(tm))
			} else {
				return domain.ThemeUser{}, fmt.Errorf("field %s: unsupported struct type %s", field.Name, fv.Type())
			}

		case reflect.Pointer:
			elemType := fv.Type().Elem()
			switch elemType {
			case reflect.TypeOf(time.Time{}):
				// Если в Redis пустая строка или специальное значение — можно трактовать как nil.
				if raw == "" {
					fv.Set(reflect.Zero(fv.Type()))
					continue
				}

				const layout = "2026-02-26T10:30:00Z"
				tm, convErr := time.Parse(layout, raw)
				if convErr != nil {
					return domain.ThemeUser{}, fmt.Errorf("field %s: parse time from %q: %w", field.Name, raw, convErr)
				}

				ptr := reflect.New(elemType)
				ptr.Elem().Set(reflect.ValueOf(tm))
				fv.Set(ptr)

			default:
				return domain.ThemeUser{}, fmt.Errorf("field %s: unsupported ptr elem type %s", field.Name, elemType)
			}

		default:
			return domain.ThemeUser{}, fmt.Errorf("field %s: unsupported kind %s", field.Name, fv.Kind())
		}
	}

	return tu, nil
}
