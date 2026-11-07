package themes_user_transport_http

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	"go.uber.org/zap"
)

func TestGetThemesUser_LimitOffset(t *testing.T) {
	service := &getThemesUserServiceStub{
		themes: []domain.ThemeUser{
			domain.ThemeUser{ID: 1},
			domain.ThemeUser{ID: 2},
			domain.ThemeUser{ID: 3},
			domain.ThemeUser{ID: 4},
		},
	}

	handler := &ThemeUserHTTPHandler{
		themeUserService: service,
	}

	ctx := core_logger.ToContext(
		context.Background(),
		&core_logger.Logger{Logger: zap.NewNop()},
	)

	req := httptest.NewRequest(
		"GET",
		"/api/v1/themes_user?limit=2&offset=1",
		nil,
	).WithContext(ctx)
	rec := httptest.NewRecorder()

	handler.GetThemesUser(rec, req)

	if rec.Code != 200 {
		t.Fatalf("expected HTTP status 200, got %d; body: %s", rec.Code, rec.Body.String())
	}

	if service.gotLimit == nil || *service.gotLimit != 2 {
		t.Fatalf("expected service to receive limit=2, got %v", service.gotLimit)
	}

	if service.gotOffset == nil || *service.gotOffset != 1 {
		t.Fatalf("expected service to receive offset=1, got %v", service.gotOffset)
	}

	var response []struct {
		ID int `json:"id"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v; body: %s", err, rec.Body.String())
	}

	if len(response) != 2 {
		t.Fatalf("expected 2 theme_user records in response, got %d", len(response))
	}

	if response[0].ID != 2 || response[1].ID != 3 {
		t.Fatalf("expected response IDs [2 3], got [%d %d]", response[0].ID, response[1].ID)
	}
}
