package main 

import (
	"log"
	"net/http"
	"github.com/spector-asael/lab2-asaeltobar/internal/routes"
) 

func main() {
	mux := routes.SetupRoute(http.NewServeMux())

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}