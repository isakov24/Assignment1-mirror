package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"cloud-assignment-1/internal/handlers"
)

var startTime = time.Now()

func main() {
	// Set the start time for uptime calculation in the status handler
	handlers.StartTime = startTime

	mux := http.NewServeMux()

	// Register endpoint prefixes with their respective handlers
	mux.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler)
	mux.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	mux.HandleFunc("/countryinfo/v1/exchange/", handlers.ExchangeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
