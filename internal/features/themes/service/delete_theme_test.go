package themes_service

import (
	"context"
	"errors"
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
)

func TestThemeService_DeleteTheme_GetIDsErrorDoesNotDeleteTheme(t *testing.T) {
	const themeID = 123

	t.Setenv("LOGGER_FOLDER", t.TempDir())

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		t.Fatalf("init application logger: %v", err)
	}
	defer logger.Close()

	ctx := core_logger.ToContext(context.Background(), logger)

	wantErr := errors.New("failed to get theme user IDs")
	theme := domain.Theme{ID: themeID}

	postgres := &postgresRepoStub{
		themes: []domain.Theme{theme},
	}

	themeUserPostgres := &themeUserCachePostgresStub{
		err: wantErr,
	}

	themeUserRedis := &themeUserCacheRedisStub{}

	themeRedis := &redisRepoStub{
		themeCache:  map[string]domain.Theme{"theme:123": theme},
		themesCache: map[string][]domain.Theme{},
		listVersion: 7,
	}

	service := &ThemeService{
		themeRepositoryPostgres: postgres,
		themeUserCachePostgres:  themeUserPostgres,
		themeUserCacheRedis:     themeUserRedis,
		themeRepositoryRedis:    themeRedis,
	}

	err = service.DeleteTheme(ctx, themeID)

	if !errors.Is(err, wantErr) {
		t.Fatalf("DeleteTheme() error = %v, want wrapped error %v", err, wantErr)
	}

	if themeUserPostgres.getIDsByThemeIDCalls != 1 {
		t.Errorf("GetIDsByThemeID calls = %d, want 1", themeUserPostgres.getIDsByThemeIDCalls)
	}

	if len(postgres.themes) != 1 || postgres.themes[0].ID != themeID {
		t.Errorf("Postgres theme was deleted; themes = %+v", postgres.themes)
	}

	if themeUserRedis.invalidateByIDsCalls != 0 {
		t.Errorf("InvalidateByIDs calls = %d, want 0", themeUserRedis.invalidateByIDsCalls)
	}

	if _, ok := themeRedis.themeCache["theme:123"]; !ok {
		t.Error("theme was deleted from Redis cache")
	}

	if themeRedis.listVersion != 7 {
		t.Errorf("themes list version = %d, want 7", themeRedis.listVersion)
	}
}
