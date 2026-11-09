package integration_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_errors "github.com/LayMan011/Golang-My-Study/internal/core/errors"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
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

func Test_Delete_User_With_Theme(t *testing.T) {
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	features := InitFeatures(ctx, TestDB, redisClient)

	user, err := features.UsersService.CreateUser(ctx, domain.User{
		ID:        42,
		Version:   1,
		Email:     "user@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов 1",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	theme, err := features.ThemesService.CreateTheme(ctx, domain.Theme{
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

	if _, err = features.UsersService.GetUser(ctx, user.ID); err != nil {
		t.Fatalf("get user before delete: %v", err)
	}

	if _, err = features.ThemesService.GetTheme(ctx, theme.ID); err != nil {
		t.Fatalf("get theme before delete: %v", err)
	}

	err = features.UsersService.DeleteUser(ctx, user.ID)

	if !errors.Is(err, core_errors.ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}

	if _, err = features.UsersRepository.GetUser(ctx, user.ID); err != nil {
		t.Fatalf("expected ErrNotFound user, got %v", err)
	}

	if _, err = features.UsersService.GetUser(ctx, user.ID); err != nil {
		t.Fatalf("expected ErrNotFound user, got %v", err)
	}

	_, err = features.ThemesRepository.GetTheme(ctx, theme.ID)
	if err != nil {
		t.Fatalf("expected ErrNotFound theme, got %v", err)
	}

	_, err = features.ThemesService.GetTheme(ctx, theme.ID)
	if err != nil {
		t.Fatalf("expected ErrNotFound theme, got %v", err)
	}
}

func Test_Delete_Theme_With_ThemeUser(t *testing.T) {
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	features := InitFeatures(ctx, TestDB, redisClient)

	user, err := features.UsersService.CreateUser(ctx, domain.User{
		ID:        42,
		Version:   1,
		Email:     "user1@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов 2",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	theme, err := features.ThemesService.CreateTheme(ctx, domain.Theme{
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

	theme_user, err := features.ThemeUserService.CreateThemeUser(
		ctx,
		domain.NewThemeUserUnitialized(theme.ID, user.ID),
	)
	if err != nil {
		t.Fatalf("create theme-user relation: %v", err)
	}

	_, err = features.ThemesService.GetTheme(ctx, theme.ID)
	if err != nil {
		t.Fatalf("get theme relation before delete: %v", err)
	}

	if _, err := features.ThemeUserService.GetThemeUser(ctx, theme_user.ID); err != nil {
		t.Fatalf("get theme_user relation before delete: %v", err)
	}

	if err := features.ThemesService.DeleteTheme(ctx, theme.ID); err != nil {
		t.Fatalf("delete theme: %v", err)
	}

	_, err = features.ThemeUserRepository.GetThemeUser(ctx, theme_user.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme_user, got %v", err)
	}

	_, err = features.ThemeUserService.GetThemeUser(ctx, theme_user.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme_user, got %v", err)
	}

	_, err = features.ThemesRepository.GetTheme(ctx, theme.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme, got %v", err)
	}

	_, err = features.ThemesService.GetTheme(ctx, theme.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme, got %v", err)
	}

	var exists bool
	err = TestDB.QueryRow(ctx, `
	SELECT EXISTS (
		SELECT 1
		FROM progress.themes_user
		WHERE id = $1
	)
`, theme_user.ID).Scan(&exists)
	if err != nil {
		t.Fatalf("check theme_user relation in Postgres: %v", err)
	}

	if exists {
		t.Fatal("expected theme_user relation to be absent from Postgres")
	}
}

func Test_Delete_User_With_ThemeUser(t *testing.T) {
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	features := InitFeatures(ctx, TestDB, redisClient)

	user_1, err := features.UsersService.CreateUser(ctx, domain.User{
		ID:        41,
		Version:   1,
		Email:     "user3@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов Иванович",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	theme, err := features.ThemesService.CreateTheme(ctx, domain.Theme{
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
		AuthorUserID:    user_1.ID,
	})
	if err != nil {
		t.Fatalf("create theme: %v", err)
	}

	user_2, err := features.UsersService.CreateUser(ctx, domain.User{
		ID:        42,
		Version:   1,
		Email:     "user4@example.com",
		CreatedAt: time.Now(),
		Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
		FullName:  "Иван Иванов Иванович",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	if user_1.ID == user_2.ID {
		t.Fatalf("expected distinct users, got same ID %d", user_1.ID)
	}

	theme_user, err := features.ThemeUserService.CreateThemeUser(
		ctx,
		domain.NewThemeUserUnitialized(theme.ID, user_2.ID),
	)
	if err != nil {
		t.Fatalf("create theme_user relation: %v", err)
	}

	if _, err := features.UsersService.GetUser(ctx, user_2.ID); err != nil {
		t.Fatalf("get user relation before delete: %v", err)
	}

	if _, err := features.ThemeUserService.GetThemeUser(ctx, theme_user.ID); err != nil {
		t.Fatalf("get theme_user relation before delete: %v", err)
	}

	if err := features.UsersService.DeleteUser(ctx, user_2.ID); err != nil {
		t.Fatalf("delete user: %v", err)
	}

	_, err = features.UsersRepository.GetUser(ctx, user_2.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound user, got %v", err)
	}

	_, err = features.UsersService.GetUser(ctx, user_2.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound user, got %v", err)
	}

	_, err = features.ThemeUserRepository.GetThemeUser(ctx, theme_user.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme_user, got %v", err)
	}

	_, err = features.ThemeUserService.GetThemeUser(ctx, theme_user.ID)
	if !errors.Is(err, core_errors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound theme_user, got %v", err)
	}
}

type Features struct {
	UsersService        *users_service.UserService
	ThemesService       *themes_service.ThemeService
	ThemeUserService    *themes_user_service.ThemeUserService
	UsersRepository     *users_postgres_repository.UserRepository
	ThemesRepository    *themes_postgres_repository.ThemeRepository
	ThemeUserRepository *themes_user_postgres_repository.ThemeUserRepository
}

func InitFeatures(
	ctx context.Context,
	pool *core_pgx_pool.Pool,
	redisClient *core_goredis_pool.RedisClient,
) *Features {
	themesUserRepositoryPostgres := themes_user_postgres_repository.NewThemeUserRepository(pool)
	themesUserRepositoryRedis := themes_user_redis_repository.NewThemesUserRepository(redisClient)
	themesUserService := themes_user_service.NewThemeUserService(themesUserRepositoryPostgres, themesUserRepositoryRedis)

	usersRepositoryPostgres := users_postgres_repository.NewUsersRepository(pool)
	usersRepositoryRedis := users_redis_repository.NewUsersRepository(redisClient)
	usersService := users_service.NewUsersService(
		usersRepositoryPostgres,
		usersRepositoryRedis,
		themesUserRepositoryPostgres,
		themesUserRepositoryRedis,
	)

	themesRepositoryPostgres := themes_postgres_repository.NewThemeRepository(pool)
	themesRepositoryRedis := themes_redis_repository.NewThemesRepository(redisClient)
	themesService := themes_service.NewThemeService(
		themesRepositoryPostgres,
		themesRepositoryRedis,
		themesUserRepositoryPostgres,
		themesUserRepositoryRedis,
	)

	return &Features{
		UsersService:        usersService,
		ThemesService:       themesService,
		ThemeUserService:    themesUserService,
		UsersRepository:     usersRepositoryPostgres,
		ThemesRepository:    themesRepositoryPostgres,
		ThemeUserRepository: themesUserRepositoryPostgres,
	}
}
