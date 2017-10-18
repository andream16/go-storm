package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"github.com/andream16/go-storm/api/feature/price/service"
	"fmt"
)

var itemHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getPrice,
	"POST"    : postPrice,
	"PUT"  	  : putPrice,
	"DELETE"  : deletePrice,
}

func PriceHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		priceHandlersMap, ok := itemHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, priceHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getPrice
// @Description gets all prices of an item given an itemId.
// @Accept  json
// @Param   item        	query   string    true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /price
// @Router /price [get]
func getPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	itemId := r.URL.Query().Get("item"); if len(itemId) == 0 {
		return response.Response{Status: "Bad Request", Message: "No item id was passed."}, "badRequest"
	}
	prices, pricesError := service.GetPrices(itemId, db); if pricesError != nil {
		return response.Response{Status: "Internal Server Error", Message: pricesError.Error()}, "serverError"
	}
	return prices, ""
}

// @Title postPrice
// @Description add [](price, date) for a given item.
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Param   prices        	query   request.Prices    true        "prices"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /price
// @Router /price [post]
func postPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var prices request.Prices
	decodePricesErr := json.NewDecoder(r.Body).Decode(&prices); if decodePricesErr != nil || len(prices.Prices) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addPricesError := service.AddPrices(prices, db); if addPricesError != nil {
		return response.Response{Status: "Bad Request", Message: addPricesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: "Successfully added Price."}, ""
}

// @Title putPrice
// @Description deletes and reinserts all the prices for a given item
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /price
// @Router /price [put]
func putPrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var prices request.Prices
	decodePricesErr := json.NewDecoder(r.Body).Decode(&prices); if decodePricesErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	putPricesError := service.EditPrice(prices, db); if putPricesError != nil {
		return response.Response{Status: "Internal Server Error", Message: putPricesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully updated prices for item %s", prices.Item)}, ""
}

// @Title putPrice
// @Description deletes all the prices for a given item
// @Accept  json
// @Param   item        	query   string    		  true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /price
// @Router /price [delete]
func deletePrice(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var prices request.Prices
	decodePricesErr := json.NewDecoder(r.Body).Decode(&prices); if decodePricesErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deletePricesError := service.DeletePrice(prices.Item, db); if deletePricesError != nil {
		return response.Response{Status: "Internal Server Error", Message: deletePricesError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully deleted prices for item %s", prices.Item)}, ""
}
