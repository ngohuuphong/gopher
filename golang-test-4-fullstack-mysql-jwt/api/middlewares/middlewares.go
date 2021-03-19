package middlewares

import (
	"context"
	"golang-test-4-fullstack-mysql-jwt/api/auth"
	"golang-test-4-fullstack-mysql-jwt/api/models"
	"golang-test-4-fullstack-mysql-jwt/api/utils/types"
	"log"
	"net/http"
)

// SetMiddlewareLogger displays a info message of the API
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// SetMiddlewareJSON set the application Content-Type
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication authorize an access
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := auth.ExtractToken(w, r)
		if token == nil {
			return
		}
		log.Println(token)
		if token.Valid {
			ctx := context.WithValue(
				r.Context(),
				types.UserKey("user"),
				token.Claims.(*models.Claim).User,
			)
			next(w, r.WithContext(ctx))
		}
	}
}
