package themes_user_postgres_repository

import core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"

type ThemeUserRepository struct {
	pool core_postgres_pool.Pool
}

func NewThemeUserRepository(
	pool core_postgres_pool.Pool,
) *ThemeUserRepository {
	return &ThemeUserRepository{
		pool: pool,
	}
}
