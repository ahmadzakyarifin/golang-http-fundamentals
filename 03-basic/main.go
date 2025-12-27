package main

import "net/http"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// cara lama 
	// id := r.URL.Query().Get("id")

	// cara baru 
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hai saya sedang ambil data ( Method GET)"))
	case http.MethodPost:
		w.Write([]byte("Hai saya sedang kirim data ( Method POST )"))
	case http.MethodPatch:
		if id == "" {
			w.Write([]byte("Eror : Maaf id belum ada"))
		}else {
			w.Write([]byte("Hai saya sedang edit data sebagian yang di miliki oleh id : " + id))
		}
	case http.MethodPut:
		if id == "" {
			w.Write([]byte("Eror : Maaf id belum ada"))
		}else {
			w.Write([]byte("Hai saya sedang edit data seluruhnya yang di miliki oleh id : " + id))
		}
	case http.MethodDelete:
		if id == "" {
			w.Write([]byte("Eror : Maaf id belum ada"))
		}else {
			w.Write([]byte("Hai saya sedang hapud data yang di miliki oleh id : " + id))
		}		
	}
} 

func main() {
	http.HandleFunc("/", apiHandler)
	http.HandleFunc("/{id}", apiHandler)
	println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
