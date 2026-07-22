package themes_user_redis_repository

import core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"

type ThemeUserRepository struct {
	pool core_redis_pool.Pool
}

func NewThemesUserRepository(
	pool core_redis_pool.Pool,
) *ThemeUserRepository {
	return &ThemeUserRepository{
		pool: pool,
	}
}
