package middleware

import (
	"context"
	"net/http"
)


type key int 
const userKey key = 0

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey).(string)
	return user, ok
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		currentUser := "zaky"

		ctx := context.WithValue(r.Context(), userKey, currentUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
