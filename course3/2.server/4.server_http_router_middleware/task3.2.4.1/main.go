package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/1", handleRoute1)
	r.Get("/2", handleRoute2)
	r.Get("/3", handleRoute3)

	http.ListenAndServe(":8080", r)
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 1"))
	w.WriteHeader(http.StatusOK)
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 2"))
	w.WriteHeader(http.StatusOK)
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 3"))
	w.WriteHeader(http.StatusOK)
}
