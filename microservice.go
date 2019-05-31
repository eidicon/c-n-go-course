package main

import (
	"fmt"
	"net/http"
	"os"
	"cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/api/v1/echo", api.Echo)
	http.HandleFunc("/api/v1/books", api.BooksHandleFunc)
	http.HandleFunc("/api/v1/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello cloud native go.")
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong!")
}

