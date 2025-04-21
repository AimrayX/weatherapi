package middleware

import (
	"net/http"
	"fmt"
)

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("🔥 CORS middleware triggered for:", r.Method, r.URL.Path)
		
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost") // Set this to your frontend URL
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, hx-request, hx-target, hx-trigger, hx-current-url")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Only if you're using cookies or credentials

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
