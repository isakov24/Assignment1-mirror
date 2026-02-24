package models

// CountryInfo represents the information about a country returned by the info endpoint
type CountryInfo struct {
	Name       string            `json:"name"`
	Capital    string            `json:"capital"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`       // ny
	Continents []string          `json:"continents"` // ny
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
}

// ExchangeInfo represents the exchange rates of neighboring countries' currencies
type ExchangeInfo struct {
	Country       string             `json:"country"`
	BaseCurrency  string             `json:"base-currency"`  // merk bindestrek
	ExchangeRates map[string]float64 `json:"exchange-rates"` // merk bindestrek
}
