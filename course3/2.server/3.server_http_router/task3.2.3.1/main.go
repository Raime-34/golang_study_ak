package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/1", firstHandler)
	r.Get("/2", secondHandler)
	r.Get("/3", thirdHandler)

	http.ListenAndServe(":8080", r)
}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
	w.WriteHeader(http.StatusOK)
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 2"))
	w.WriteHeader(http.StatusOK)
}

func thirdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 3"))
	w.WriteHeader(http.StatusOK)
}
