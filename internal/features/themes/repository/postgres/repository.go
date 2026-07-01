package themes_postgres_repository

import core_postgres_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool"

type ThemeRepository struct {
	pool core_postgres_pool.Pool
}

func NewThemeRepository(
	pool core_postgres_pool.Pool,
) *ThemeRepository {
	return &ThemeRepository{
		pool: pool,
	}
}
