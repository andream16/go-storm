package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/api/feature/item/service"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"strconv"
)

var itemHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getItem,
	"POST"    : postItem,
	"PUT"  	  : putItem,
	"DELETE"  : deleteItem,
}

func ItemHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

// @Title getItem
// @Description gets an item given its id as url param, else, checks for page and sizes and returns a slice of size items
// @Accept  json
// @Param   item        	query   string    false        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /item [get]
func getItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	page := r.URL.Query().Get("page"); if len(page) == 0 {
		itemId := r.URL.Query().Get("item"); if len(itemId) == 0 {
			return response.Response{}, "badRequest"
		}
		item := service.GetItem(itemId, db)
		if len(item.Item) == 0 || item.Item == ""  {
			return response.Response{Status: "Not Found", Message: "Item not found"}, "badRequest"
		}
		return item, ""
	} else {
		size := r.URL.Query().Get("size"); if len(size) == 0 {
			return response.Response{Status: "Not Found", Message: "Bad request"}, "badRequest"
		}
		p, pError := strconv.Atoi(page); if pError != nil {
			return response.Response{Status: "Bad Request", Message: "Bad value for page"}, "badRequest"
		}
		s, sError := strconv.Atoi(size); if sError != nil {
			return response.Response{Status: "Bad Request", Message: "Bad value for size"}, "badRequest"
		}
		items, itemsError := service.GetItems(p, s, db); if itemsError != nil {
			return response.Response{Status: "Not Found", Message: itemsError.Error()}, "badRequest"
		}
		return request.Items{items}, ""
	}
}

// @Title postItem
// @Description adds an item given its body or updates if already exists
// @Accept  json
// @Param   item        	query   string	      true        "item"
// @Param   manufacturer    query   string 		  true        "manufacturer"
// @Param   url             query   string 		  false       "url"
// @Param   image           query   string 		  false       "image"
// @Param   title           query   string 		  false       "title"
// @Param   description     query   string 		  false       "description"
// @Param   has_reviews     query   bool 		  false       "has_reviews"
// @Success 200 {object} response.Response    {Status :"Ok", Message: "Successfully inserted item"}
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /item [post]
func postItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var item request.Item
	decodeErr := json.NewDecoder(r.Body).Decode(&item); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	insertError := service.AddItem(item, db); if insertError != nil {
		return response.Response{Status: "Internal Server Error", Message: insertError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: "Successfully inserted item"}, ""
}

// @Title putItem
// @Description updates an item given its id
// @Accept  json
// @Param   item        	query   string	      true        "item"
// @Param   manufacturer    query   string 		  false       "manufacturer"
// @Param   url             query   string 		  false       "url"
// @Param   image           query   string 		  false       "image"
// @Param   title           query   string 		  false       "title"
// @Param   description     query   string 		  false       "description"
// @Param   has_reviews     query   bool 		  false       "has_reviews"
// @Success 200 {object} response.Response    {Status :"Ok", Message: "Successfully updated item"}
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /item [put]
func putItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var item request.Item
	decodeErr := json.NewDecoder(r.Body).Decode(&item); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	updateError := service.EditItem(item, db); if updateError != nil {
		return response.Response{Status: "Internal Server Error", Message: updateError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: "Successfully updated item"}, ""
}

// @Title deleteItem
// @Description deletes an item given its body
// @Accept  json
// @Param   item        	query   string	      true        "item"
// @Param   manufacturer    query   string 		  false       "manufacturer"
// @Param   url             query   string 		  false       "url"
// @Param   image           query   string 		  false       "image"
// @Param   title           query   string 		  false       "title"
// @Param   description     query   string 		  false       "description"
// @Param   has_reviews     query   bool 		  false       "has_reviews"
// @Success 200 {object} response.Response    {Status :"Ok", Message: "Successfully deleted item"}
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /item [delete]
func deleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var item request.Item
	decodeErr := json.NewDecoder(r.Body).Decode(&item); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteError := service.DeleteItem(item.Item, db); if deleteError != nil {
		return response.Response{Status: "Internal Server Error", Message: deleteError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: "Successfully deleted item"}, ""
}
