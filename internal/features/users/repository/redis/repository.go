package users_redis_repository

import core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"

type UserRepository struct {
	pool core_redis_pool.Pool
}

func NewUsersRepository(
	pool core_redis_pool.Pool,
) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
