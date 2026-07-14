package users_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
)

func (r *UserRepository) GetToken(ctx context.Context, login string) (*domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	if err := r.pool.DoError(ctx, "SELECT", 0); err != nil {
		return nil, fmt.Errorf("db switch error: %w", err)
	}

	key := fmt.Sprintf("jwt:%s", login)

	jsonData, err := r.pool.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var pair domain.TokenPair
	err = json.Unmarshal([]byte(jsonData), &pair)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &pair, nil
}
