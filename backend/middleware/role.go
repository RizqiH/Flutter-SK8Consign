package middleware

import (
	"encoding/json"
	"net/http"
)

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-Role")

		if role != "admin" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "Access denied. Admin role required",
			})
			return
		}

		next(w, r)
	})
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(next)
}



