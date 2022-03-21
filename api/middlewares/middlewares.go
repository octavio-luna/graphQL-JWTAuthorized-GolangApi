package middlewares

import (
	"errors"
	"net/http"

	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/auth"
	responses "github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/utils/response_handlers"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.WriteError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
