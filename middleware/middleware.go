package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dinel13/lanjukang/pkg/utilities"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// if r.Method == "OPTIONS" {
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}

func ChecToken(w http.ResponseWriter, r *http.Request, jwtSecret string) (int, int, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return 0, 0, errors.New("invalid token")
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	id, role, err := utilities.ParseToken(tokenString, jwtSecret)
	if err != nil {
		return 0, 0, err

	}
	if id == 0 {
		return 0, 0, errors.New("invalid token")

	}

	return id, role, nil
}
