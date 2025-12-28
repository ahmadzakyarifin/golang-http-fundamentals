package handler

import (
	"fmt"
	"net/http"
)

func Barang(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Saya mau ambil barang"))
}

func CreateBarang(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Saya mau buat data barang baru "))	
}

func UpdateSebagianBarang(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pesan := fmt.Sprintf("Saya mau edit beberapa bagian yang di miliki oleh barang ber-id %s" ,id)
	w.Write([]byte(pesan))
}

func UpdateSemuaBarang(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pesan := fmt.Sprintf("Saya mau edit seluruh bagian yang di miliki oleh barang ber-id %s" ,id)
	w.Write([]byte(pesan))	
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pesan := fmt.Sprintf("Saya mau hapus barang ber-id %s" ,id)
	w.Write([]byte(pesan))	
}
