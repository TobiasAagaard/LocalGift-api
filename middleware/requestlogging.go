package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		timeStarted := time.Now()

		responseWrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(responseWrapped, r)

		duration := time.Since(timeStarted)

		log.Printf(
			"%s %s %d %dms",
			r.Method,
			r.RequestURI,
			responseWrapped.statusCode,
			duration.Milliseconds(),
		)
	}))
}
