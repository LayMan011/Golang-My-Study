package themes_user_redis_repository

import (
	"context"
	"fmt"
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

	parseInt := func(redisKey string) (int, error) {
		raw, ok := data[redisKey]
		if !ok {
			return 0, nil
		}

		n, err := strconv.Atoi(raw)
		if err != nil {
			return 0, fmt.Errorf("field %s: parse int from %q: %w", redisKey, raw, err)
		}

		return n, nil
	}

	if tu.ID, err = parseInt("id"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.Version, err = parseInt("version"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.Percentages, err = parseInt("percentages"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.TotalLessons, err = parseInt("total_lessons"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.CompletedLessons, err = parseInt("completed_lessons"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.ThemeID, err = parseInt("theme_id"); err != nil {
		return domain.ThemeUser{}, err
	}

	if tu.UserID, err = parseInt("user_id"); err != nil {
		return domain.ThemeUser{}, err
	}

	if raw, ok := data["completed"]; ok {
		tu.Completed, err = strconv.ParseBool(raw)
		if err != nil {
			return domain.ThemeUser{}, fmt.Errorf("field Completed: parse bool from %q: %w", raw, err)
		}
	}

	if raw, ok := data["addition_at"]; ok {
		tu.AdditionAt, err = time.Parse(time.RFC3339Nano, raw)
		if err != nil {
			return domain.ThemeUser{}, fmt.Errorf("field AdditionAt: parse time from %q: %w", raw, err)
		}
	}

	if raw, ok := data["completed_at"]; ok && raw != "" {
		completedAt, parseErr := time.Parse(time.RFC3339Nano, raw)
		if parseErr != nil {
			return domain.ThemeUser{}, fmt.Errorf("field CompletedAt: parse time from %q: %w", raw, parseErr)
		}

		tu.CompletedAt = &completedAt
	}

	return tu, nil
}
