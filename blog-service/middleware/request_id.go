package middleware

import (
	"blog-service/constants"
	"blog-service/logger"

	"context"
	"net/http"

	"github.com/google/uuid"
)

// RequestIDMiddleware is an HTTP middleware that generates a unique request ID for each incoming request.
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := getMiddlewareRequestID(r)

		// Add the request ID to the request context
		updatedCtx := context.WithValue(r.Context(), constants.RequestIDKey, requestID)

		// Log the request ID for tracking purposes
		logger.Log.Info(updatedCtx, "Received request: %s %s", r.Method, r.URL.Path)

		// Pass the updated context to the next handler in the chain
		r = r.WithContext(updatedCtx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the completion of the request
		logger.Log.Info(updatedCtx, "Completed request: %s %s", r.Method, r.URL.Path)
	})
}

func getMiddlewareRequestID(r *http.Request) string {
	// Retrieve the request ID from the header or generate a new one
	requestID := r.Header.Get(constants.HeaderRequestID)
	if requestID == "" {
		requestID = uuid.New().String()
	}
	return requestID
}
