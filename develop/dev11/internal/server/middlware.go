package server

import (
	"log"
	"net/http"
)

func logMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Запрос:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
