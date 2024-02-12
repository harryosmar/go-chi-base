package middlewares

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func ResponseSetHeaderRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to application/json
		w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(r.Context()))

		// Call the next handlers
		next.ServeHTTP(w, r)
	})
}
