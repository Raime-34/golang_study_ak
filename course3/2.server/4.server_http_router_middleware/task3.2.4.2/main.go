package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(LoggerMiddleware)
	r.Get("/", handleRoute)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Request received",
			// Structured context as strongly typed Field values.
			zap.String("url", r.URL.String()),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		next.ServeHTTP(w, r)
	})
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
	w.WriteHeader(http.StatusOK)
}
