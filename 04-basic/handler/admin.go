package handler

import "net/http"


func Dasboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tampilakan semua grafik dan persentase"))

}