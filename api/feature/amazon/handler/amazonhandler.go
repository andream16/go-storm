package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/model/request"
	"encoding/json"
	"fmt"
	"github.com/andream16/go-storm/api/feature/amazon/service"
)

var amazonHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"POST"    : postAmazon,
}


func AmazonHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reviewHandlersMap, ok := amazonHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, reviewHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title postAmazon
// @Description adds one amazon entry .
// @Accept  json
// @Param   amazon       query  request.Amazon    true        "amazon"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /amazon
// @Router /amazon [post]
func postAmazon(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var amazonEntry request.Amazon
	decodeAmazonErr := json.NewDecoder(r.Body).Decode(&amazonEntry); if decodeAmazonErr != nil {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addAmazonError := service.AddAmazonEntry(amazonEntry, db); if addAmazonError != nil {
		return response.Response{Status: "Bad Request", Message: addAmazonError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully added amazon entry for item %s.", amazonEntry.Item)}, ""
}
