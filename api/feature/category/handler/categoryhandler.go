package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/category/service"
	"github.com/andream16/go-storm/model/request"
	"encoding/json"
	"fmt"
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

// @Title getCategory
// @Description gets categories given an itemId
// @Accept  json
// @Param   item        	query   string	      true        "item"
// @Param   has_reviews     query   bool 		  false       "has_reviews"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /category [get]
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

// @Title postCategory
// @Description adds categories for a given item
// @Accept  json
// @Param   item        	query   request.CategoryRequest  true
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /category [post]
func postCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var categoriesByItem request.CategoryRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&categoriesByItem); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addCategoriesByItemError := service.AddCategoriesByItem(categoriesByItem, db); if addCategoriesByItemError != nil {
		return response.Response{Status: "Internal Server Error", Message: addCategoriesByItemError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: fmt.Sprintf("Successfully categories for item %s", categoriesByItem.Item)}, ""
}

// @Title putCategory
// @Description edits categories for a given item
// @Accept  json
// @Param   item        	query   request.CategoryRequest  true
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /category [put]
func putCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var categoriesByItem request.CategoryRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&categoriesByItem); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	editCategoryError := service.EditCategory(categoriesByItem, db); if editCategoryError != nil {
		return response.Response{Status: "Internal Server Error", Message: editCategoryError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: fmt.Sprintf("Successfully edited categories for item %s", categoriesByItem.Item)}, ""
}

// @Title deleteCategory
// @Description deletes categories for a given item
// @Accept  json
// @Param   item        	query   request.CategoryRequest  true
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /category [put]
func deleteCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var categoriesByItem request.CategoryRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&categoriesByItem); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteCategoryError := service.DeleteCategory(categoriesByItem, db); if deleteCategoryError != nil {
		return response.Response{Status: "Internal Server Error", Message: deleteCategoryError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: fmt.Sprintf("Successfully deleted categories for item %s", categoriesByItem.Item)}, ""
}