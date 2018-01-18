package handler

import (
	"net/http"
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/functionmapper"
	"github.com/andream16/go-storm/shared/handler/errortostatus"
	"github.com/andream16/go-storm/shared/handler/response"
	"github.com/andream16/go-storm/api/feature/review_tmp/service"
	"encoding/json"
	"github.com/andream16/go-storm/model/request"
	"fmt"
)

var reviewTmpHandlers = map[string]func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, string) {
	"POST"    : postTmpReview,
}


func ReviewTmpHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reviewTmpHandlersMap, ok := reviewTmpHandlers[r.Method]; if ok {
			res, errorMessage := functionmapper.FunctionMapper(w, r, db, reviewTmpHandlersMap); if errorMessage != "" {
				errortostatus.ErrorAsStringToStatus(errorMessage, w)
			}
			response.JsonResponse(res, w)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func postTmpReview(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string) {
	var review request.ReviewTmp
	decodeReviewTmpErr := json.NewDecoder(r.Body).Decode(&review); if decodeReviewTmpErr != nil || len(review.Reviews) == 0  {
		return response.Response{Status: "Bad Request", Message: "Bad body"}, "badRequest"
	}
	addTmpReviewError := service.AddTmpReview(review, db); if addTmpReviewError != nil {
		return response.Response{Status: "Bad Request", Message: addTmpReviewError.Error()}, "serverError"
	}
	return response.Response{Status: "Ok", Message: fmt.Sprintf("Successfully added tmp review for item %s.", review.Item)}, ""
}