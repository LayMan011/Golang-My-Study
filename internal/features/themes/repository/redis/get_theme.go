package themes_redis_repository

import (
	"context"
	"fmt"
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

	if t.ID, err = parseInt("id"); err != nil {
		return domain.Theme{}, err
	}

	if t.Version, err = parseInt("version"); err != nil {
		return domain.Theme{}, err
	}

	if t.AllRatings, err = parseInt("all_ratings"); err != nil {
		return domain.Theme{}, err
	}

	if t.NumberOfRatings, err = parseInt("number_of_ratings"); err != nil {
		return domain.Theme{}, err
	}

	if t.NumberOfUsers, err = parseInt("number_of_users"); err != nil {
		return domain.Theme{}, err
	}

	if t.Price, err = parseInt("price"); err != nil {
		return domain.Theme{}, err
	}

	if t.AuthorUserID, err = parseInt("author_user_id"); err != nil {
		return domain.Theme{}, err
	}

	t.Title = data["title"]
	t.Subject = data["subject"]
	t.Level = data["level"]
	t.Duration = data["duration"]
	t.Format = data["format"]

	if raw, ok := data["description"]; ok {
		t.Description = &raw
	}

	if raw, ok := data["rating"]; ok {
		rating, parseErr := strconv.ParseFloat(raw, 64)
		if parseErr != nil {
			return domain.Theme{}, fmt.Errorf("field Rating: parse float from %q: %w", raw, parseErr)
		}

		t.Rating = &rating
	}

	if raw, ok := data["created_at"]; ok {
		t.CreatedAt, err = time.Parse(time.RFC3339Nano, raw)
		if err != nil {
			return domain.Theme{}, fmt.Errorf("field CreatedAt: parse time from %q: %w", raw, err)
		}
	}

	return t, nil
}
