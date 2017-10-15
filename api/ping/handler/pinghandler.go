package handler

import (
	"net/http"
	"github.com/andream16/go-storm/shared/response"
)

func PingHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		response.JsonResponse(response.Response{"Ok", "Pong"}, w)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
