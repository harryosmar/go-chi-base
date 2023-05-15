package middlewares

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"net/http"
)

func CustomRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(middleware.RequestIDHeader)
		if requestID == "" {
			guid := uuid.New()
			requestID = guid.String()
		}
		ctx = context.WithValue(ctx, middleware.RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
