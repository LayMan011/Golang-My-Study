package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LayMan011/Golang-My-Study/internal/core/domain"
	core_logger "github.com/LayMan011/Golang-My-Study/internal/core/logger"
	core_http_request "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/request"
	core_http_response "github.com/LayMan011/Golang-My-Study/internal/core/transport/http/response"
)

type GetstatisticsResponse struct {
	ThemesCreated               int      `json:"themes_created"`
	ThemesCompleted             int      `json:"themes_completed"`
	ThemesCompletedRate         *float64 `json:"themes_completed_rate"`
	ThemesAverageCompletionTime *string  `json:"themes_average_completion_time"`
}

func toDTOFromDomain(statistics domain.Statistics) GetstatisticsResponse {
	var avgTime *string
	if statistics.ThemesAverageCompletionTime != nil {
		duration := statistics.ThemesAverageCompletionTime.String()
		avgTime = &duration
	}

	return GetstatisticsResponse{
		ThemesCreated:               statistics.ThemesCreated,
		ThemesCompleted:             statistics.ThemesCompleted,
		ThemesCompletedRate:         statistics.ThemesCompletedRate,
		ThemesAverageCompletionTime: avgTime,
	}
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/from/to query params",
		)

		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)

		return
	}

	response := toDTOFromDomain(statistics)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	usesrID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return usesrID, from, to, nil
}
