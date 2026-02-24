package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud-assignment-1/internal/models"
)

// RestCountry represents the structure of the country data returned by the REST Countries API
type RestCountry struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	Capital []string `json:"capital"`

	Population int `json:"population"`

	Area float64 `json:"area"`

	Continents []string `json:"continents"`

	Languages map[string]string `json:"languages"`

	Borders []string `json:"borders"`

	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`

	Flags struct {
		Png string `json:"png"`
	} `json:"flags"`
}

// GetCountry fetches country information from the REST Countries API based on the provided country code
func GetCountry(code string) (*RestCountry, error) {
	url := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("restcountries returned status %d", resp.StatusCode)
	}

	var data []RestCountry
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no country data found")
	}

	return &data[0], nil
}

// Converts the raw RestCountry data into the public CountryInfo model used by the API response
func ConvertToCountryInfo(rc *RestCountry) *models.CountryInfo {
	capital := ""
	if len(rc.Capital) > 0 {
		capital = rc.Capital[0]
	}

	return &models.CountryInfo{
		Name:       rc.Name.Common,
		Capital:    capital,
		Population: rc.Population,
		Area:       rc.Area,
		Continents: rc.Continents,
		Languages:  rc.Languages,
		Borders:    rc.Borders,
		Flag:       rc.Flags.Png,
	}
}

// Checks the status of the REST Countries API by making a simple request and returning the HTTP status code
func CheckRestCountriesAPI() int {
	resp, err := http.Get("http://129.241.150.113:8080/v3.1/all")
	if err != nil {
		return 0
	}
	return resp.StatusCode
}
