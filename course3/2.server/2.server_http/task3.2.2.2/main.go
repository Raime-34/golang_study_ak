package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var port string

	err := godotenv.Load()
	if err != nil {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello world"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(port, nil)
}
