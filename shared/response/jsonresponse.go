package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status	string `json:"status"`
	Message string `json:"message"`
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}