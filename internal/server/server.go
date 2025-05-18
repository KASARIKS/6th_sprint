package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type routerData struct {
	Logger *log.Logger
	Server http.Server
}

func CreateServer(logger *log.Logger) *routerData {
	registerHandlers()

	router := newRouter(logger)
	return router
}

func newRouter(logger *log.Logger) *routerData {
	return &routerData{
		Logger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      http.DefaultServeMux,
			ErrorLog:     logger,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 10,
			IdleTimeout:  time.Second * 15,
		},
	}
}

func registerHandlers() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
}
