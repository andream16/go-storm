package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"github.com/andream16/go-storm/api/feature/currency/service"
	"fmt"
)
var itemHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getCurrency,
	"POST"    : postCurrency,
	"PUT"  	  : putCurrency,
	"DELETE"  : deleteCurrency,
}

func CurrencyHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		currencyHandlersMap, ok := itemHandlers[r.Method];
		if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, currencyHandlersMap);
			if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getCurrency
// @Description gets all currency entries of a currency given a currency name.
// @Accept  json
// @Param   name        	query   string    true        "name"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /currency
// @Router /currency [get]
func getCurrency(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	name := r.URL.Query().Get("name"); if len(name) == 0 {
		return response.Response{Status: "Bad Request", Message: "No currency name was passed."}, "badRequest"
	}
	currencies, currenciesError := service.GetCurrencyByName(name, db); if currenciesError != nil {
		return response.Response{Status: "Internal Server Error", Message: currenciesError.Error()}, "serverError"
	}
	return currencies, ""
}

// @Title postCurrency
// @Description add [](name,value, date) for a given currency.
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Param   currencies       query   request.Currency true        "currencies"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /currency
// @Router /currency [post]
func postCurrency(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var currencies request.Currency
	decodeCurrenciesErr := json.NewDecoder(r.Body).Decode(&currencies); if decodeCurrenciesErr != nil || len(currencies.Trend) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addCurrenciesError := service.AddCurrencies(currencies, db); if addCurrenciesError != nil {
		return response.Response{Status: "Bad Request", Message: addCurrenciesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: "Successfully added currencies."}, ""
}


// @Title putCurrency
// @Description deletes and reinserts all the currencies for a given currency name
// @Accept  json
// @Param   name        	query   string    		  true        "name"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /currency
// @Router /currency [put]
func putCurrency(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var currencies request.Currency
	decodeCurrenciesErr := json.NewDecoder(r.Body).Decode(&currencies); if decodeCurrenciesErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	putCurrenciesError := service.EditCurrency(currencies, db); if putCurrenciesError != nil {
		return response.Response{Status: "Internal Server Error", Message: putCurrenciesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully updated currencies for currency %s", currencies.Name)}, ""
}

// @Title deleteCurrency
// @Description deletes all the currencies for a given currency name
// @Accept  json
// @Param   name        	query   string    		  true        "name"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /currency
// @Router /currency [delete]
func deleteCurrency(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var currencies request.Currency
	decodeCurrenciesErr := json.NewDecoder(r.Body).Decode(&currencies); if decodeCurrenciesErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteCurrenciesError := service.DeleteCurrency(currencies.Name, db); if deleteCurrenciesError != nil {
		return response.Response{Status: "Internal Server Error", Message: deleteCurrenciesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully deleted currencies for currency %s", currencies.Name)}, ""
}