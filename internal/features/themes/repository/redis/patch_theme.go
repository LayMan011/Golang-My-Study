package themes_redis_repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *ThemeRepository) PatchTheme(
	ctx context.Context,
	id int,
	patch domain.ThemePatch,
) error {
	data := make(map[string]string)
	fillStringPatchTheme(data, patch)

	if len(data) == 0 {
		return nil
	}

	key := fmt.Sprintf("theme:%d", id)

	if err := r.pool.HSet(ctx, key, data); err != nil {
		return fmt.Errorf("patch theme in redis: %w", err)
	}

	return nil
}

func fillStringPatchTheme(data map[string]string, patch domain.ThemePatch) {
	v := reflect.ValueOf(patch)
	t := v.Type()

	fieldToKey := map[string]string{
		"Title":       "title",
		"Description": "description",
		"Subject":     "subject",
		"Level":       "level",
		"Duration":    "duration",
		"Format":      "format",
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		key, ok := fieldToKey[field.Name]
		if !ok {
			continue
		}

		fv := v.Field(i) // это Nullable[string]

		setField := fv.FieldByName("Set")
		valueField := fv.FieldByName("Value")

		if !setField.IsValid() || !valueField.IsValid() {
			continue
		}

		if setField.Bool() && !valueField.IsNil() {
			str := valueField.Elem().String()
			data[key] = str
		}
	}
}
