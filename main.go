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
	mux.HandleFunc("POST /reciepts/process", apiCfg.handlerProcessReciepts) // Reciept  // Return ID

	// Determine and return points awarded to a reciept (GET)
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID) // ID  // Return points

	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on port: %v", port)
	log.Fatal(srv.ListenAndServe())

}
