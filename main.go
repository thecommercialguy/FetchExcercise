package main

import (
	"log"
	"net/http"
	"sync"
)

type apiConfig struct {
	DB sync.Map
}

func main() {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	mux := http.NewServeMux()

	// Processes and stores receipts (POST)
	mux.HandleFunc("POST /receipts/process", apiCfg.handlerProcessReceipts) // Receipt  // Return ID

	// Determines and returns points awarded to a receipt (GET)
	mux.HandleFunc("GET /receipts/{id}/points", apiCfg.handlerGetPointsByID) // ID  // Return points

	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on port: %v", port)
	log.Fatal(srv.ListenAndServe())

}
