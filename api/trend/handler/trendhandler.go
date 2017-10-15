package handler

import (
	"net/http"
	"database/sql"
)

func TrendHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
