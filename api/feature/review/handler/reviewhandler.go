package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/review/service"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"fmt"
)

var reviewHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"GET"     : getReview,
	"POST"    : postReview,
	"PUT"  	  : putReview,
	"DELETE"  : deleteReview,
}


func ReviewHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reviewHandlersMap, ok := reviewHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, reviewHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// @Title getReview
// @Description gets all reviews given a item.
// @Accept  json
// @Param   item        	query   string    true        "item"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /trend
// @Router /trend [get]
func getReview(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	item := r.URL.Query().Get("item"); if len(item) == 0 {
		return response.Response{Status: "Not Found", Message: "Bad item entry."}, "badRequest"
	}
	review, reviewError := service.GetReviewByItem(item, db); if reviewError != nil {
		return response.Response{Status: "Not Found", Message: reviewError.Error()}, "badRequest"
	}
	return review, ""
}

// @Title postReview
// @Description adds n review entries given a item.
// @Accept  json
// @Param   review       query  request.Review    true        "review"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /review
// @Router /review [post]
func postReview(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var review request.Review
	decodeReviewErr := json.NewDecoder(r.Body).Decode(&review); if decodeReviewErr != nil || len(review.Reviews) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addReviewError := service.AddReview(review, db); if addReviewError != nil {
		return response.Response{Status: "Bad Request", Message: addReviewError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully added review for item %s.", review.Item)}, ""
}

// @Title putReview
// @Description edits n review entries given an item.
// @Accept  json
// @Param   review        	query  request.Review    true        "review"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /review
// @Router /review [put]
func putReview(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var review request.Review
	decodeReviewErr := json.NewDecoder(r.Body).Decode(&review); if decodeReviewErr != nil || len(review.Reviews) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addReviewError := service.EditReviewByItem(review, db); if addReviewError != nil {
		return response.Response{Status: "Bad Request", Message: addReviewError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully edited review for item %s.", review.Item)}, ""
}

// @Title deleteReview
// @Description deletes n review entries given an item.
// @Accept  json
// @Param   review        	query  request.Review    true        "review"
// @Success 200 {object} response.Response    response.Response
// @Failure 403 {object} response.Response    {Status: "Bad Request", Message: error.Error()}
// @Failure 500 {object} response.Response    {Status: "Internal Server Error", error.Error()}
// @Resource /review
// @Router /review [delete]
func deleteReview(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var review request.Review
	decodeReviewErr := json.NewDecoder(r.Body).Decode(&review); if decodeReviewErr != nil || len(review.Reviews) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	deleteReviewError := service.DeleteReviewByItem(review, db); if deleteReviewError != nil {
		return response.Response{Status: "Bad Request", Message: deleteReviewError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully deleted review for item %s.", review.Item)}, ""
}
