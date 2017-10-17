package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"github.com/andream16/go-storm/api/price/service"
)

var itemHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getPrice,
	"POST"    : postPrice,
	"PUT"  	  : putPrice,
	"DELETE"  : deletePrice,
}

func PriceHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		itemHandlersMap, ok := itemHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, itemHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return response.Response{}, ""
}

func postPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var price request.Price
	var prices []request.Price
	decodePriceErr := json.NewDecoder(r.Body).Decode(&price); if decodePriceErr != nil {
		decodePricesErr := json.NewDecoder(r.Body).Decode(&prices); if decodePricesErr != nil {
			return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
		}
		addPricesError := service.AddPrices(prices, db); if addPricesError != nil {
			return response.Response{Status: "Bad Request", Message: addPricesError.Error()}, "serverError"
		}
	}
	addPriceError := service.AddPrice(price, db); if addPriceError != nil {
		return response.Response{Status: "Bad Request", Message: addPriceError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: "Successfully added Price."}, ""
}

func putPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return response.Response{}, ""
}
func deletePrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return response.Response{}, ""
}
