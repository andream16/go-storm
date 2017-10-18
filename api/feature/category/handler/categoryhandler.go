package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
)

var categoryHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getCategory,
	"POST"    : postCategory,
	"PUT"  	  : putCategory,
	"DELETE"  : deleteCategory,
}

func CategoryHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		itemHandlersMap, ok := categoryHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, itemHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
func getCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return "", ""
}

func postCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return "", ""
}

func putCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return "", ""
}

func deleteCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	return "", ""
}