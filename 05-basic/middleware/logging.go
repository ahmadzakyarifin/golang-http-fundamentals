package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("before => method : %s , path : %s",r.Method,r.URL.Path)
		fmt.Println()
		next.ServeHTTP(w,r)
		fmt.Println("after")
	})
}