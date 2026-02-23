package middleware

import (
	"net/http"
	"log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s URL: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler 
		log.Println("Request processed")
	})
}

func WrongPathErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested path is valid
		if r.URL.Path == "/about" || r.URL.Path == "/contact" || r.URL.Path == "/quote" || r.URL.Path == "/" {
			next.ServeHTTP(w, r) // Call the next handler for valid paths
			return
		} else {
			w.Write([]byte("404 page not found\n")) // Return 404 for invalid paths
			log.Printf("Invalid Path Accessed")
			return
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}