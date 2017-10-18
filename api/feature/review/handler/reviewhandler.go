package handler

import (
	"net/http"
	"database/sql"
)

func ReviewHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}