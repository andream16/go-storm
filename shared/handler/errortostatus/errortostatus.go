package errortostatus

import "net/http"

func ErrorAsStringToStatus(err error, w http.ResponseWriter) (http.ResponseWriter) {
	switch err.Error() {
	case "serverError":
		w.WriteHeader(http.StatusInternalServerError)
		break
	case "badRequest":
		w.WriteHeader(http.StatusBadRequest)
		break
	case "forbidden":
		w.WriteHeader(http.StatusForbidden)
		break
	case "notFound":
		w.WriteHeader(http.StatusNotFound)
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
		break
	}
	return w
}
