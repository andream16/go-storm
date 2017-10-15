package handler

import (
	"net/http"
	"database/sql"
)

func PriceHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
