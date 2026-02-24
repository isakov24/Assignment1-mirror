package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"cloud-assignment-1/internal/services"
)

var StartTime time.Time

type StatusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrenciesAPI    int    `json:"currenciesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Check the status of the REST Countries API and the Currency API
	restStatus := services.CheckRestCountriesAPI()
	currencyStatus := services.CheckCurrencyAPI()

	// Calculate server uptime in seconds
	uptime := time.Since(StartTime).Seconds()

	// Build the status response
	response := StatusResponse{
		RestCountriesAPI: restStatus,
		CurrenciesAPI:    currencyStatus,
		Version:          "v1",
		Uptime:           int64(uptime),
	}
	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
