package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/manufacturer/service"
	"github.com/andream16/go-storm/model/request"
	"encoding/json"
	"fmt"
	"strconv"
)

var manufacturerHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getManufacturer,
	"POST"    : postManufacturer,
	"DELETE"  : deleteManufacturer,
}

func ManufacturerHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		manufacturerHandlersMap, ok := manufacturerHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, manufacturerHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getManufacturer
// @Description gets manufacturer given an itemId, else, checks for page and sizes and returns a slice of size manufacturers
// @Accept  json
// @Param   item        	query   string	      true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /manufacturer [get]
func getManufacturer(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
		page := r.URL.Query().Get("page"); if len(page) == 0 {
		itemId := r.URL.Query().Get("item");
		if len(itemId) == 0 {
			manufacturer := r.URL.Query().Get("manufacturer");
			if len(manufacturer) == 0 {
				return response.Response{Status: "Bad Request", Message: "Bad request, no manufacturer or item found."}, "badRequest"
			}
			items, itemsError := service.GetItemsByManufacturer(manufacturer, db);
			if itemsError != nil {
				return response.Response{Status: "Internal Server Error", Message: itemsError.Error()}, "serverError"
			}
			return items, ""
		}
		manufacturer, manufacturerError := service.GetManufacturerByItem(itemId, db);
		if manufacturerError != nil {
			return response.Response{Status: "Internal Server Error", Message: manufacturerError.Error()}, "serverError"
		}
		return manufacturer, ""
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
		manufacturers, manufacturersError := service.GetManufacturers(p, s, db); if manufacturersError != nil {
			return response.Response{Status: "Not Found", Message: manufacturersError.Error()}, "badRequest"
		}
		return request.Manufacturers{manufacturers}, ""
	}

}

// @Title postManufacturer
// @Description add manufacturer
// @Accept  json
// @Param   item        	query   request.Manufacturer  true
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /manufacturer [post]
func postManufacturer(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var manufacturer request.Manufacturer
	decodeErr := json.NewDecoder(r.Body).Decode(&manufacturer); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addManufacturerError := service.AddManufacturer(manufacturer, db); if addManufacturerError != nil {
		return response.Response{Status: "Internal Server Error", Message: addManufacturerError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: fmt.Sprintf("Successfully added manufacturer")}, ""
}

// @Title deleteManufacturer
// @Description deletes manufacturer given its name
// @Accept  json
// @Param   manufacturer        	query   request.Manufacturer  true
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /item
// @Router /manufacturer [delete]
func deleteManufacturer(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var manufacturer request.Manufacturer
	decodeErr := json.NewDecoder(r.Body).Decode(&manufacturer); if decodeErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteManufacturerError := service.DeleteManufacturer(manufacturer, db); if deleteManufacturerError != nil {
		return response.Response{Status: "Internal Server Error", Message: deleteManufacturerError.Error()}, "serverError"
	}
	return response.Response{Status :"Ok", Message: fmt.Sprintf("Successfully deleted manufacturer for name %s", manufacturer.Name)}, ""
}