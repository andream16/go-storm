package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"github.com/andream16/go-storm/api/feature/forecast/service"
	"fmt"
)

var forecastHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getForecast,
	"POST"    : postForecast,
	"PUT"  	  : putForecast,
	"DELETE"  : deleteForecast,
}

func ForecastHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		forecastHandlersMap, ok := forecastHandlers[r.Method];
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

// @Title getForecast
// @Description gets all forecasts of an item given an itemId.
// @Accept  json
// @Param   item        	query   string    true        "item"
// @Param   test_size       query   string    true        "test_size"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /forecast
// @Router /forecast [get]
func getForecast(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	itemId := r.URL.Query().Get("item"); if len(itemId) == 0 {
		return response.Response{Status: "Bad Request", Message: "No item id was passed."}, "badRequest"
	}
	forecastTestSize := r.URL.Query().Get("test_size"); if len(forecastTestSize) == 0 {
		return response.Response{Status: "Bad Request", Message: "No forecast test size found."}, "badRequest"
	}
	forecast, forecastError := service.GetForecastByItemAndForecastTestSize(itemId, forecastTestSize, db); if forecastError != nil {
		return response.Response{Status: "Internal Server Error", Message: forecastError.Error()}, "serverError"
	}
	return forecast, ""
}

// @Title postForecast
// @Description add [](price, date) for a given item.
// @Accept  json
// @Param   forecasts       query   request.Forecasts true        "forecasts"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /forecast
// @Router /forecast [post]
func postForecast(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var forecasts request.Forecast
	decodeForecastsErr := json.NewDecoder(r.Body).Decode(&forecasts); if decodeForecastsErr != nil || len(forecasts.Forecast) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addForecastsError := service.AddForecasts(forecasts, db); if addForecastsError != nil {
		return response.Response{Status: "Bad Request", Message: addForecastsError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: "Successfully added Forecast."}, ""
}


// @Title putForecast
// @Description deletes and reinserts all the forecasts for a given item
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /forecast
// @Router /forecast [put]
func putForecast(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var forecasts request.Forecast
	decodeForecastsErr := json.NewDecoder(r.Body).Decode(&forecasts); if decodeForecastsErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	putForecastsError := service.EditForecast(forecasts, db); if putForecastsError != nil {
		return response.Response{Status: "Internal Server Error", Message: putForecastsError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully updated forecasts for item %s", forecasts.Item)}, ""
}

// @Title deleteForecast
// @Description deletes all the forecasts for a given item
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /forecast
// @Router /forecast [delete]
func deleteForecast(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var forecasts request.Forecast
	decodeForecastsErr := json.NewDecoder(r.Body).Decode(&forecasts); if decodeForecastsErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteForecastsError := service.DeleteForecast(forecasts.Item, forecasts.TestSize, db); if deleteForecastsError != nil {
		return response.Response{Status: "Internal Server Error", Message: deleteForecastsError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully deleted forecasts for item %s", forecasts.Item)}, ""
}