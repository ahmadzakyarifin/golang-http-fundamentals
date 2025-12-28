package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadzakyarifin/04-basic/handler"
)

func main(){
	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /barang",handler.Barang)

	mux.HandleFunc("POST /barang/create",handler.CreateBarang)

	mux.HandleFunc("PATCH /barang/{id}",handler.UpdateSebagianBarang)

	mux.HandleFunc("PUT /barang/{id}",handler.UpdateSemuaBarang)

	mux.HandleFunc("DELETE /barang/{id}",handler.DeleteBarang)

	mux.HandleFunc("GET /dasboard",handler.Dasboard)
	
	go func() {
		admin := &http.Server{
		Addr: ":9090",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 20 * time.Second,
		}
		fmt.Println("http://localhost:9090")
	    admin.ListenAndServe()	
	}()

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	fmt.Println("http://localhost:8080")
	server.ListenAndServe()

}