package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

// gzipResponseWriter wraps http.ResponseWriter to write gzip responses
type gzipResponeWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponeWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

func CompressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipResponeWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(w, r)
		fmt.Println("Sent response from compression middleware")
	})
}
