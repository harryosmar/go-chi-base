package middlewares

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidateRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			next.ServeHTTP(w, r)
			return
		}

		validate := validator.New()
		err := validate.Struct(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Request body is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
