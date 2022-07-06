package handlers

import (
	"log"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
