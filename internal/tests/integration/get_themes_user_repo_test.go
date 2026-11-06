package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_pgx_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool/pgx"
	core_goredis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool/goredis"
	themes_user_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/theme_users/repository/postgres"
	themes_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/themes/repository/postgres"
	users_postgres_repository "github.com/LayMan011/Golang-My-Study/internal/features/users/repository/postgres"
)

func Test_Delete_User_With_Theme(t *testing.T) {
	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	usersRepo, themesRepo, themesUserRepo := initFeatures(ctx, TestDB, redisClient)

	if err := generateUsers(ctx, usersRepo, 1); err != nil {
		t.Fatalf("failed to generate themes: %v", err)
	}

	themes_id, err := generateThemes(ctx, themesRepo, 1, 1)
	if err != nil {
		t.Fatalf("failed to generate themes: %v", err)
	}

	if err := generateThemeUsers(ctx, themesUserRepo, 1, themes_id[:(len(themes_id)/2)]); err != nil {
		t.Fatalf("failed to generate themes_user: %v", err)
	}

	if err := generateThemeUsers(ctx, themesUserRepo, 2, themes_id[(len(themes_id)/2):]); err != nil {
		t.Fatalf("failed to generate themes_user: %v", err)
	}

	user_id := 2

	themes_user_user_id_2, err := themesUserRepo.GetThemesUser(ctx, &user_id, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_user_id_2); got != 5 {
		t.Fatalf("ожидалось 5 записей пользователя %d, получено %d", user_id, got)
	}

	for _, item := range themes_user_user_id_2 {
		if item.UserID != user_id {
			t.Errorf("получена запись чужого пользователя: ожидался user_id=%d, получен %d", user_id, item.UserID)
		}
	}

	theme_id := 2

	themes_user_theme_id_2, err := themesUserRepo.GetThemesUser(ctx, nil, &theme_id, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_theme_id_2); got != 1 {
		t.Fatalf("ожидалось 1 записей пользователя %d, получено %d", user_id, got)
	}

	for _, item := range themes_user_theme_id_2 {
		if item.ThemeID != theme_id {
			t.Errorf("получена запись другой темы: ожидался theme_id=%d, получен %d", theme_id, item.ThemeID)
		}
	}

	limit := 2

	themes_user_with_limit, err := themesUserRepo.GetThemesUser(ctx, nil, nil, &limit, nil)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_with_limit); got != 2 {
		t.Fatalf("ожидалось 2 записей из-за limit, получено %d", got)
	}

	limit_zero := 0

	themes_user_with_limit_zero, err := themesUserRepo.GetThemesUser(ctx, nil, nil, &limit_zero, nil)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_with_limit_zero); got != 0 {
		t.Fatalf("ожидалось 0 записей из-за limit_zero, получено %d", got)
	}

	offset := 1

	themes_user_with_offset, err := themesUserRepo.GetThemesUser(ctx, nil, nil, nil, &offset)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if themes_user_with_limit[1] != themes_user_with_offset[0] {
		t.Fatalf(
			"не совпала запись на границе страниц: limit вернул theme_id=%d, offset — theme_id=%d",
			themes_user_with_limit[1].ThemeID,
			themes_user_with_offset[0].ThemeID,
		)
	}

	themes_user_with_limit_and_offset, err := themesUserRepo.GetThemesUser(ctx, nil, nil, &limit, &offset)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_with_limit_and_offset); got != 2 {
		t.Fatalf("ожидалось 2 записей из-за limit_zero, получено %d", got)
	}

	if themes_user_with_limit[1] != themes_user_with_limit_and_offset[0] {
		t.Fatalf(
			"не совпала запись на границе страниц: limit вернул theme_id=%d, offset — theme_id=%d",
			themes_user_with_limit[1].ThemeID,
			themes_user_with_limit_and_offset[0].ThemeID,
		)
	}

	user_id_not_exists := 20

	themes_user_empty, err := themesUserRepo.GetThemesUser(ctx, &user_id_not_exists, nil, nil, nil)
	if err != nil {
		t.Fatalf("GetThemesUser: %v", err)
	}

	if got := len(themes_user_empty); got != 0 {
		t.Fatalf("ожидалось 0 записей из-за не существющего user_id, получено %d", got)
	}
}

func generateThemes(
	ctx context.Context,
	themesRepo *themes_postgres_repository.ThemeRepository,
	user_id int,
	r int,
) ([]int, error) {
	themes_id := make([]int, 0, 10)

	for i := r; i < r+10; i++ {
		theme, err := themesRepo.CreateTheme(ctx, domain.Theme{
			ID:              i,
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
			AuthorUserID:    user_id,
		})
		if err != nil {
			return []int{}, fmt.Errorf("create theme #%d: %w", i, err)
		}

		themes_id = append(themes_id, theme.ID)
	}

	return themes_id, nil
}

func generateUsers(
	ctx context.Context,
	usersRepo *users_postgres_repository.UserRepository,
	r int,
) error {
	for i := r; i < r+10; i++ {
		_, err := usersRepo.CreateUser(ctx, domain.User{
			ID:        i,
			Version:   1,
			Email:     fmt.Sprintf("user%d@example.com", i),
			CreatedAt: time.Now(),
			Password:  []byte("c29tZV9wYXNzd29yZF9oYXNo"),
			FullName:  "Иван Иванов 1",
		})
		if err != nil {
			return fmt.Errorf("create user #%d: %w", i, err)
		}
	}

	return nil
}

func generateThemeUsers(
	ctx context.Context,
	themesUserRepo *themes_user_postgres_repository.ThemeUserRepository,
	user_id int,
	themes_id []int,
) error {
	for i, theme_id := range themes_id {
		_, err := themesUserRepo.CreateThemeUser(
			ctx,
			domain.NewThemeUserUnitialized(theme_id, user_id),
		)
		if err != nil {
			return fmt.Errorf("create theme_user #%d: %w", i, err)
		}
	}

	return nil
}

func initFeatures(
	ctx context.Context,
	pool *core_pgx_pool.Pool,
	redisClient *core_goredis_pool.RedisClient,
) (*users_postgres_repository.UserRepository, *themes_postgres_repository.ThemeRepository, *themes_user_postgres_repository.ThemeUserRepository) {
	usersRepositoryPostgres := users_postgres_repository.NewUsersRepository(pool)

	themesRepositoryPostgres := themes_postgres_repository.NewThemeRepository(pool)

	themesUserRepository := themes_user_postgres_repository.NewThemeUserRepository(pool)

	return usersRepositoryPostgres, themesRepositoryPostgres, themesUserRepository
}
