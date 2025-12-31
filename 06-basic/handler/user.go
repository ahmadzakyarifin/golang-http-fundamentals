package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmadzakyarifin/golang-http-fundamentals/06-basic/middleware"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Halo, %s 👋", user)
}
