package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"cloud-assignment-1/internal/models"
	"cloud-assignment-1/internal/services"
)

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract country code from URL path
	path := r.URL.Path
	prefix := "/countryinfo/v1/exchange/"
	if !strings.HasPrefix(path, prefix) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	code := strings.TrimPrefix(path, prefix)
	if code == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	// Fetch country info from REST Countries API
	country, err := services.GetCountry(code)
	if err != nil {
		http.Error(w, "Country not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Determine base currency for the country
	var baseCurrency string
	for k := range country.Currencies {
		baseCurrency = k
		break
	}
	if baseCurrency == "" {
		http.Error(w, "No currency found for this country", http.StatusInternalServerError)
		return
	}

	// Fetch neighboring countries (borders)
	borders := country.Borders
	if len(borders) == 0 {
		response := models.ExchangeInfo{
			Country:       country.Name.Common,
			BaseCurrency:  baseCurrency,
			ExchangeRates: map[string]float64{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// For each neighboring country, fetch its currency
	type neighborResult struct {
		currencies []string
	}

	neighborChan := make(chan neighborResult, len(borders))
	var wg sync.WaitGroup

	for _, borderCode := range borders {
		wg.Add(1)
		go func(bc string) {
			defer wg.Done()
			neighbor, err := services.GetCountry(bc)
			if err != nil {
				neighborChan <- neighborResult{currencies: []string{}}
				return
			}
			var curList []string
			for c := range neighbor.Currencies {
				curList = append(curList, c)
			}
			neighborChan <- neighborResult{currencies: curList}
		}(borderCode)
	}

	go func() {
		wg.Wait()
		close(neighborChan)
	}()

	targetCurrencies := make(map[string]bool)
	for res := range neighborChan {
		for _, cur := range res.currencies {
			targetCurrencies[cur] = true
		}
	}

	// Fetch exchange rates for the base currency from the currency API
	allRates, err := services.GetCurrencyRates(baseCurrency)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract exchange rates for target currencies
	exchangeRates := make(map[string]float64)
	for target := range targetCurrencies {
		if rate, exists := allRates.Rates[target]; exists {
			exchangeRates[target] = rate
		}
	}

	// Build and return the response
	response := models.ExchangeInfo{
		Country:       country.Name.Common,
		BaseCurrency:  baseCurrency,
		ExchangeRates: exchangeRates,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
