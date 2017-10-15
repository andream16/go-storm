package functionmapper

import (
	"database/sql"
	"github.com/andream16/go-storm/shared/handler/response"
	"net/http"
)

func FunctionMapper(w http.ResponseWriter, r *http.Request, db *sql.DB,
	f func(w http.ResponseWriter, r *http.Request, db *sql.DB) (response.Response, error)) (response.Response, error) {
	return f(w, r, db)
}
