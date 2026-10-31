package integration_test

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"testing"
	"time"

	core_pgx_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/postgres/pool/pgx"
	core_goredis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool/goredis"
	"github.com/golang-migrate/migrate/v4"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	testcontainersredis "github.com/testcontainers/testcontainers-go/modules/redis"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	TestDB      *core_pgx_pool.Pool
	redisClient *core_goredis_pool.RedisClient
)

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось запустить PostgreSQL: %v\n", err)
		return 1
	}
	defer terminatePostgres(postgresContainer)

	postgresHost, err := postgresContainer.Host(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось получить хост PostgreSQL: %v\n", err)
		return 1
	}

	postgresPort, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось получить порт PostgreSQL: %v\n", err)
		return 1
	}

	setupCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	TestDB, err = core_pgx_pool.NewPool(setupCtx, core_pgx_pool.Config{
		Host:     postgresHost,
		Port:     postgresPort.Port(),
		User:     "test",
		Password: "test",
		Database: "test",
		Timeout:  5 * time.Second,
	})
	cancel()

	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось создать пул PostgreSQL: %v\n", err)
		return 1
	}
	defer TestDB.Close()

	migrationURL := url.URL{
		Scheme: "file",
		Path:   "../../../migrations",
	}

	postgresURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("test", "test"),
		Host:   net.JoinHostPort(postgresHost, postgresPort.Port()),
		Path:   "test",
	}
	query := postgresURL.Query()
	query.Set("sslmode", "disable")
	postgresURL.RawQuery = query.Encode()

	migration, err := migrate.New(migrationURL.String(), postgresURL.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось инициализировать миграции: %v\n", err)
		return 1
	}
	defer migration.Close()

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Fprintf(os.Stderr, "не удалось применить миграции: %v\n", err)
		return 1
	}

	redisContainer, err := testcontainersredis.Run(ctx, "redis:7-alpine")
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось запустить Redis: %v\n", err)
		return 1
	}
	defer terminateRedis(redisContainer)

	redisHost, err := redisContainer.Host(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось получить хост Redis: %v\n", err)
		return 1
	}

	redisPort, err := redisContainer.MappedPort(ctx, "6379/tcp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось получить порт Redis: %v\n", err)
		return 1
	}

	redisConfig := core_goredis_pool.RedisConfig{
		ADDR:     net.JoinHostPort(redisHost, redisPort.Port()),
		Password: "",
		Database: 0,
		Timeout:  5 * time.Second,
	}

	redisClient = core_goredis_pool.CreateRedisClientMust(redisConfig)
	defer redisClient.Close()

	return m.Run()
}

func terminatePostgres(container *postgres.PostgresContainer) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := container.Terminate(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "не удалось остановить PostgreSQL: %v\n", err)
	}
}

func terminateRedis(container *testcontainersredis.RedisContainer) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := container.Terminate(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "не удалось остановить Redis: %v\n", err)
	}
}
