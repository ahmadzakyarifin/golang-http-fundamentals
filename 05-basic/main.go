package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadzakyarifin/golang-http-fundamentals/05-basic/middleware"
)

func main(){
	mux := http.NewServeMux()

	handler:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word"))
	})

	handlerMid := middleware.LoggingMiddleware(handler)

	mux.Handle("/",handlerMid)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}