package http

import (
	"log"
	"net/http"
	"time"
	"github.com/go-chi/cors"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func CORSAll() func(http.Handler) http.Handler {
	return cors.AllowAll().Handler
}
