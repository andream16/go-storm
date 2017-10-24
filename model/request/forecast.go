package request

type Forecast struct {
	Item string `json:"item,omitempty"`
	Name string   `json:"name,omitempty"`
	Forecast []ForecastEntry
}

type ForecastEntry struct {
	Price float64 `json:"price,omitempty"`
	Date string`json:"date,omitempty"`
}

