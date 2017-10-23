package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/category/service"
)

var categoryHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getCategory,
	"POST"    : postCategory,
	"PUT"  	  : putCategory,
	"DELETE"  : deleteCategory,
}

func CategoryHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryHandlersMap, ok := categoryHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, categoryHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	itemId := r.URL.Query().Get("item"); if len(itemId) == 0 {
		category := r.URL.Query().Get("category"); if len(category) == 0 {
			return response.Response{Status: "Bad Request", Message: "Bad request, no category or item found."}, "badRequest"
		}
		items, itemsError := service.GetItemsByCategory(category, db); if itemsError != nil {
			return response.Response{Status: "Internal Server Error", Message: itemsError.Error()}, "serverError"
		}
		return items, ""
	}
	categories, categoriesError := service.GetCategoriesByItem(itemId, db); if categoriesError != nil {
		return response.Response{Status: "Internal Server Error", Message: categoriesError.Error()}, "serverError"
	}
	return categories, ""
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