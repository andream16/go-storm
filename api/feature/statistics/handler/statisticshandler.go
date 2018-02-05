package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"github.com/andream16/go-storm/api/feature/statistics/service"
)

var statisticsHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getStatistics,
	"POST"    : postStatistics,
}

func StatisticsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		forecastHandlersMap, ok := statisticsHandlers[r.Method];
		if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, forecastHandlersMap);
			if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getStatistics
// @Description gets all forecast statistics given an Item ID and a test size.
// @Accept  json
// @Param   item        	query   string    true        "item"
// @Param   test_size       query   string    true        "test_size"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /statistics
// @Router /statistics [get]
func getStatistics(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	itemId := r.URL.Query().Get("item"); if len(itemId) == 0 {
		return response.Response{Status: "Bad Request", Message: "No item id was passed."}, "badRequest"
	}
	statisticsTestSize := r.URL.Query().Get("test_size"); if len(statisticsTestSize) == 0 {
		return response.Response{Status: "Bad Request", Message: "No forecast test size found."}, "badRequest"
	}
	statistics, statisticsError := service.GetStatisticsByItemAndForecastTestSize(itemId, statisticsTestSize, db); if statisticsError != nil {
		return response.Response{Status: "Internal Server Error", Message: statisticsError.Error()}, "serverError"
	}
	return statistics, ""
}

// @Title postStatistics
// @Description add [](price, date) and a test size for a given item.
// @Accept  json
// @Param   forecasts    query   request.Forecasts true        "statistics"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /statistics
// @Router /statistics [post]
func postStatistics(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var statistics request.Statistics
	decodeForecastsErr := json.NewDecoder(r.Body).Decode(&statistics); if decodeForecastsErr != nil || len(statistics.Forecast) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addForecastsError := service.AddStatistics(statistics, db); if addForecastsError != nil {
		return response.Response{Status: "Bad Request", Message: addForecastsError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: "Successfully added Statistics."}, ""
}
