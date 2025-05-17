package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
