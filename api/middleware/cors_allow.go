package middleware

import (
	"net/http"
)

// A middleware that validates sessions
// for any handlers that require any sort of session validation
func AllowCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with the appropriate origin(s)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.WriteHeader(http.StatusOK)
			return
		}

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with the appropriate origin(s)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// if no errs, and its valid, launch handler
		next.ServeHTTP(w, r)
	})
}
