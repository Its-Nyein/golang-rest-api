package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/execs", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create custom server with TLS config
	server := &http.Server{
		Addr:      port,
		Handler:   nil, // use default handler,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
