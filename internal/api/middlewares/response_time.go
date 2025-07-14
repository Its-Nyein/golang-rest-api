package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request in response time middleware")
		start := time.Now()

		// create custom response writer to capture status code
		wrappedWriter := &ResponseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start)
		w.Header().Set("X-response-time", duration.String())
		fmt.Printf("Method: %s, Path: %s, Status: %d, Duration: %s\n", r.Method, r.URL.Path, wrappedWriter.status, duration.String())
		fmt.Println("Sent response from response time middleware")
	})
}
