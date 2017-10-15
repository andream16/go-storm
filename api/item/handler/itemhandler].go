package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
)

var itemHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (response.Response, error) {
	"GET"     : getItem,
	"POST"    : postItem,
	"PUT"  	  : putItem,
	"DELETE"  : deleteItem,
}

func ItemHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		itemHandlersMap, ok := itemHandlers[r.Method]; if ok {
			res, err := functionmapper.FunctionMapper(w, r, db, itemHandlersMap); if err != nil {
				errortostatus.ErrorAsStringToStatus(err, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (response.Response, error) {
	return response.Response{Status :"Ok", Message: "GET"}, nil
}
func postItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (response.Response, error) {
	return response.Response{Status :"Ok", Message: "POST"}, nil
}
func putItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (response.Response, error) {
	return response.Response{Status :"Ok", Message: "PUT"}, nil
}
func deleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (response.Response, error) {
	return response.Response{Status :"Ok", Message: "DELETE"}, nil
}
