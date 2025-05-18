package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Logger{}

	router := server.CreateServer(&logger)

	err := http.ListenAndServe(router.Server.Addr, router.Server.Handler)
	if err != nil {
		router.Logger.Fatal(err)
	}
}
