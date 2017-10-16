package functionmapper

import (
	"database/sql"
	"net/http"
)

func FunctionMapper(w http.ResponseWriter, r *http.Request, db *sql.DB,
	f func(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, string)) (interface{}, string) {
	return f(w, r, db)
}
