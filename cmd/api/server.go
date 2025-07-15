package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/v2/internal/api/middlewares"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
	w.Write([]byte("Welcome to the API server!"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on techers route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on techers route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on techers route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on techers route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on techers route"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on students route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on students route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on students route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on students route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on students route"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route"))
	case http.MethodPost:
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("Name:", r.URL.Query().Get("name"))
		fmt.Println("Name:", r.URL.Query().Get("name"))

		// parse form data
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("Form of data:", r.Form)

		w.Write([]byte("Hello POST method on execs route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on execs route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on execs route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on execs route"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers", teachersHandler)
	mux.HandleFunc("/students", studentsHandler)
	mux.HandleFunc("/execs", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:              true,
		CheckBody:               true,
		CheckBodyOnlyForContent: "application/x-www-urlenconded",
		Whitelist:               []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	// serverMux := mw.Cors(rl.RateLimiterMiddleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.CompressionMiddleware(mw.HPP(hppOptions)(mux))))))
	serverMux := ApplyMiddlewares(mux, mw.HPP(hppOptions), mw.CompressionMiddleware, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.RateLimiterMiddleware, mw.Cors)

	// create custom server with TLS config
	server := &http.Server{
		Addr:      port,
		Handler:   serverMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
