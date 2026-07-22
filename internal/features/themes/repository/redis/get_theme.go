package themes_redis_repository

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
)

func (r *ThemeRepository) GetTheme(ctx context.Context, key string) (domain.Theme, error) {
	data, err := r.pool.HGetAll(ctx, key)
	if err != nil {
		return domain.Theme{}, fmt.Errorf("get theme from redis: %w", err)
	}

	if len(data) == 0 {
		return domain.Theme{}, core_errors.ErrNotFound
	}

	var t domain.Theme
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

		switch fv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, convErr := strconv.ParseInt(raw, 10, 64)
			if convErr != nil {
				return domain.Theme{}, fmt.Errorf("field %s: parse int from %q: %w", field.Name, raw, convErr)
			}
			fv.SetInt(n)

		case reflect.String:
			fv.SetString(raw)

		case reflect.Pointer:
			// *string, *float64 и другие ptr-типы
			elemKind := fv.Type().Elem().Kind()

			switch elemKind {
			case reflect.String:
				s := raw
				fv.Set(reflect.ValueOf(&s))

			case reflect.Float32, reflect.Float64:
				f, convErr := strconv.ParseFloat(raw, 64)
				if convErr != nil {
					return domain.Theme{}, fmt.Errorf("field %s: parse float from %q: %w", field.Name, raw, convErr)
				}
				ptr := reflect.New(fv.Type().Elem())
				ptr.Elem().SetFloat(f)
				fv.Set(ptr)

			default:
				return domain.Theme{}, fmt.Errorf("field %s: unsupported ptr elem kind %s", field.Name, elemKind)
			}

		case reflect.Struct:
			if fv.Type() == reflect.TypeOf(time.Time{}) {
				// ВАЖНО: layout должен совпадать с тем, как ты сохраняешь CreatedAt в Redis при SaveTheme.
				const layout = "2026-02-26T10:30:00Z"
				tm, convErr := time.Parse(layout, raw)
				if convErr != nil {
					return domain.Theme{}, fmt.Errorf("field %s: parse time from %q: %w", field.Name, raw, convErr)
				}
				fv.Set(reflect.ValueOf(tm))
			} else {
				return domain.Theme{}, fmt.Errorf("field %s: unsupported struct type %s", field.Name, fv.Type())
			}

		default:
			return domain.Theme{}, fmt.Errorf("field %s: unsupported kind %s", field.Name, fv.Kind())
		}
	}

	return t, nil
}
