package handlers

import (
	"calendar-server/internal/errors"
	"net/http"
)


func ValidateMethod(r *http.Request, method string) error {
	if r.Method != method {
		return errors.MethodNotAllowed
	}

	return nil
}

func ValidateQuery(w http.ResponseWriter, r *http.Request, query ...string) bool {
	for _, q := range query {
		if !r.URL.Query().Has(q) {
			w.WriteHeader(http.StatusServiceUnavailable)
			return false
		}
	}

	return true
}