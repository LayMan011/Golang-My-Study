package themes_redis_repository

import core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"

type ThemeRepository struct {
	pool core_redis_pool.Pool
}

func NewThemesRepository(
	pool core_redis_pool.Pool,
) *ThemeRepository {
	return &ThemeRepository{
		pool: pool,
	}
}
