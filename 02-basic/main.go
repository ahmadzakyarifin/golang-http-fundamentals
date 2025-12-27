package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	nama := r.URL.Query().Get("nama")

	if nama == "" {
	w.Write([]byte("Hello"))
	} else {
		pesan := fmt.Sprintf("Hello, %s",nama)
		w.Write([]byte(pesan))
	}
}

func main(){
	http.HandleFunc("/",hello)

	
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}

