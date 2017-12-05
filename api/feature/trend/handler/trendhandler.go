package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/trend/service"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"fmt"
)

var trendHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getTrend,
	"POST"    : postTrend,
	"PUT"  	  : putTrend,
	"DELETE"  : deleteTrend,
}


func TrendHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		trendHandlersMap, ok := trendHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, trendHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getTrend
// @Description gets all trends given a manufacturer.
// @Accept  json
// @Param   item        	query   string    true        "manufacturer"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /trend
// @Router /trend [get]
func getTrend(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	manufacturer := r.URL.Query().Get("manufacturer"); if len(manufacturer) == 0 {
		return response.Response{Status: "Not Found", Message: "Bad manufacturer entry."}, "badRequest"
	}
	trend, trendError := service.GetTrendByManufacturer(manufacturer, db); if trendError != nil {
		return response.Response{Status: "Not Found", Message: trendError.Error()}, "badRequest"
	}
	return trend, ""
}

// @Title postTrend
// @Description adds n trend entries given a manufacturer.
// @Accept  json
// @Param   trend        	query  request.Trend    true        "trend"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /trend
// @Router /trend [post]
func postTrend(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var trend request.Trend
	decodeTrendErr := json.NewDecoder(r.Body).Decode(&trend); if decodeTrendErr != nil || len(trend.Trend) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addTrendError := service.AddTrendByManufacturer(trend, db); if addTrendError != nil {
		return response.Response{Status: "Bad Request", Message: addTrendError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully added Trend for manufacturer %s.", trend.Manufacturer)}, ""
}

// @Title putTrend
// @Description edits n trend entries given a manufacturer.
// @Accept  json
// @Param   trend        	query  request.Trend    true        "trend"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /trend
// @Router /trend [put]
func putTrend(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var trend request.Trend
	decodeTrendErr := json.NewDecoder(r.Body).Decode(&trend); if decodeTrendErr != nil || len(trend.Trend) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addTrendError := service.EditTrendByManufacturer(trend, db); if addTrendError != nil {
		return response.Response{Status: "Bad Request", Message: addTrendError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully edited Trend for manufacturer %s.", trend.Manufacturer)}, ""
}

// @Title deleteTrend
// @Description deletes n trend entries given a manufacturer.
// @Accept  json
// @Param   trend        	query  request.Trend    true        "trend"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /trend
// @Router /trend [delete]
func deleteTrend(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var trend request.Trend
	decodeTrendErr := json.NewDecoder(r.Body).Decode(&trend); if decodeTrendErr != nil || len(trend.Trend) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteTrendError := service.DeleteTrendByManufacturer(trend, db); if deleteTrendError != nil {
		return response.Response{Status: "Bad Request", Message: deleteTrendError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully deleted Trend for manufacturer %s.", trend.Manufacturer)}, ""
}
