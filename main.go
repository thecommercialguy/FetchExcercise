package main

import (
	"log"
	"net/http"
	"sync"
	// "encoding/json"
	// "time"
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

	// Will "process" the reciept(determine points)
	mux.HandleFunc("POST /reciepts/process", apiCfg.handlerProcessReciepts) // Reciept  // Return ID

	// Will get the reciept's points
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID) // ID  // Return points
	// mux.HandleFunc()

	srv := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Printf("Serving files on port: 5000")
	log.Fatal(srv.ListenAndServe())

}
