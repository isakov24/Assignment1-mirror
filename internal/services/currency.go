package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CurrencyAPIResponse represents the structure of the response from the Currency API
type CurrencyAPIResponse struct {
	Result             string             `json:"result"`
	Provider           string             `json:"provider"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
	TimeEolUnix        int64              `json:"time_eol_unix"`
	BaseCode           string             `json:"base_code"`
	Rates              map[string]float64 `json:"rates"`
}

// Fetches exchange rates for a given base currency code from the Currency API
func GetCurrencyRates(baseCode string) (*CurrencyAPIResponse, error) {
	url := fmt.Sprintf("http://129.241.150.113:9090/currency/%s", baseCode)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call currency API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("currency API returned status %d for base %s", resp.StatusCode, baseCode)
	}

	var data CurrencyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode currency API response: %w", err)
	}

	if data.Result != "success" {
		return nil, fmt.Errorf("currency API returned result: %s", data.Result)
	}

	return &data, nil
}

// Fetches exchange rates for a given base currency and a list of target currencies
// Returns a map of target currency codes to their exchange rates relative to the base currency
func GetSpecificRates(baseCode string, targets []string) (map[string]float64, error) {
	if len(targets) == 0 {
		return make(map[string]float64), nil
	}

	// Fetch all exchange rates
	allRates, err := GetCurrencyRates(baseCode)
	if err != nil {
		return nil, err
	}

	// Filter the rates to include only the target currencies
	filtered := make(map[string]float64)
	for _, target := range targets {
		if rate, exists := allRates.Rates[target]; exists {
			filtered[target] = rate
		}
		// If one is missing, we can choose to ignore it or return an error. Here we ignore missing targets
	}

	return filtered, nil
}

// Checks the status of the Currency API by making a simple request and returning the HTTP status code
func CheckCurrencyAPI() int {
	resp, err := http.Get("http://129.241.150.113:9090/currency/USD")
	if err != nil {
		return 0 // 0 indicates an error occurred while trying to reach the API
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
