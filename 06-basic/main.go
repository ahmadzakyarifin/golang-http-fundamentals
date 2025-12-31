package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ahmadzakyarifin/golang-http-fundamentals/06-basic/handler"
	"github.com/ahmadzakyarifin/golang-http-fundamentals/06-basic/middleware"
)

func main() {
	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(handler.UserHandler)

	handlerWithMiddleware := middleware.TimeoutMiddleware(
		middleware.AuthMiddleware(userHandler),
	)

	mux.Handle("/", handlerWithMiddleware)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
