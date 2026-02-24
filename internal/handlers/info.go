package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"cloud-assignment-1/internal/services"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the country code from URL path
	path := r.URL.Path
	prefix := "/countryinfo/v1/info/"
	if !strings.HasPrefix(path, prefix) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	code := strings.TrimPrefix(path, prefix)
	if code == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	// Fetch country info from the REST Countries API
	country, err := services.GetCountry(code)
	if err != nil {
		http.Error(w, "Country not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Convert the raw API model into the public CountryInfo response model
	info := services.ConvertToCountryInfo(country)

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
