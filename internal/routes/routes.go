package routes

import (
	"net/http"
	"github.com/spector-asael/lab2-asaeltobar/internal/handlers"
	"github.com/spector-asael/lab2-asaeltobar/internal/middleware"
)

func SetupRoute(mux *http.ServeMux) http.Handler{ // Function to set up the HTTP request multiplexer. 
	// Uses handler functions from the handlers package, along with middleware from the middleware fle.
	mux.HandleFunc("/", handlers.HomeHandler) 
	mux.HandleFunc("/about", handlers.AboutHandler) 
	mux.HandleFunc("/contact", handlers.ContactHandler)
	mux.HandleFunc("/quote", handlers.QuoteHandler)
	WrongPathErrorMiddleware := middleware.WrongPathErrorMiddleware(mux)  // handle wrong paths after logging.
	loggedMiddleware := middleware.LoggingMiddleware(WrongPathErrorMiddleware) // log the requests first.
	return loggedMiddleware
}