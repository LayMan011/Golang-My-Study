package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_pgx_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool/pgx"
	core_goredis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool/goredis"
	themes_user_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/theme_users/repository/postgres"
	themes_user_redis_repository "github.com/LayMan011/Golang-My-Study/internal/features/theme_users/repository/redis"
	themes_user_service "github.com/LayMan011/Golang-My-Study/internal/features/theme_users/service"
	themes_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/themes/repository/postgres"
	themes_redis_repository "github.com/LayMan011/Golang-My-Study/internal/features/themes/repository/redis"
	themes_service "github.com/LayMan011/Golang-My-Study/internal/features/themes/service"
	users_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/users/repository/postgres"
	users_redis_repository "github.com/LayMan011/Golang-My-Study/internal/features/users/repository/redis"
	users_service "github.com/LayMan011/Golang-My-Study/internal/features/users/service"
)

func Test_Delete_User_And_Theme_With_ThemeUser(t *testing.T) {
	ctx := context.Background()

	usersService, themesService, themesUserService :=
		InitFeatures(ctx, TestDB, redisClient)

	user, err := usersService.CreateUser(ctx, domain.User{
		ID:        42,
		Version:   1,
		Email:     "user@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов Иванович",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	theme, err := themesService.CreateTheme(ctx, domain.Theme{
		ID:              21,
		Version:         1,
		Title:           "123",
		Description:     nil,
		CreatedAt:       time.Now(),
		Subject:         "Русский язык",
		Rating:          nil,
		AllRatings:      0,
		NumberOfRatings: 0,
		NumberOfUsers:   0,
		Price:           5000,
		Level:           "beginner",
		Duration:        "3 месяца",
		Format:          "video",
		AuthorUserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("create theme: %v", err)
	}

	theme_user, err := themesUserService.CreateThemeUser(
		ctx,
		domain.NewThemeUserUnitialized(theme.ID, user.ID),
	)
	if err != nil {
		t.Fatalf("create theme-user relation: %v", err)
	}

	if err := themesService.DeleteTheme(ctx, theme.ID); err != nil {
		t.Fatalf("delete theme: %v", err)
	}

	_, err = themesUserService.GetThemeUser(ctx, theme_user.ID)
	if err == nil {
		t.Fatal("expected theme-user relation to be deleted with the theme")
	}

	// Пока что закомментировал эту проврку, потому что здесь все ещё есть проверка из #23. Как залью в main #23,
	// так сразу раскомментирую
	// if err := usersService.DeleteUser(ctx, user.ID); err == nil {
	// 	t.Fatal("expected deleting user to fail, got nil")
	// }

	// if _, err = usersService.GetUser(ctx, user.ID); err != nil {
	// 	t.Fatal("expected user to be deleted")
	// }
}

func InitFeatures(
	ctx context.Context,
	pool *core_pgx_pool.Pool,
	redisClient *core_goredis_pool.RedisClient,
) (*users_service.UserService, *themes_service.ThemeService, *themes_user_service.ThemeUserService) {
	usersRepositoryPostgres := users_postgres_repository.NewUsersRepository(pool)
	usersRepositoryRedis := users_redis_repository.NewUsersRepository(redisClient)
	usersService := users_service.NewUsersService(usersRepositoryPostgres, usersRepositoryRedis)

	themesRepositoryPostgres := themes_postgres_repository.NewThemeRepository(pool)
	themesRepositoryRedis := themes_redis_repository.NewThemesRepository(redisClient)
	themesService := themes_service.NewThemeService(themesRepositoryPostgres, themesRepositoryRedis)

	themesUserRepository := themes_user_postgres_repository.NewThemeUserRepository(pool)
	themesUserRepositoryRedis := themes_user_redis_repository.NewThemesUserRepository(redisClient)
	themesUserService := themes_user_service.NewThemeUserService(themesUserRepository, themesUserRepositoryRedis)

	return usersService, themesService, themesUserService
}
