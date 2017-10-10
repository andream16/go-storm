package request

import "time"

type Forecast struct {
	Item string `json:"item"`
	Forecast []ForecastEntry
}

type ForecastEntry struct {
	Price float64 `json:"price"`
	Date time.Time `json:"date"`
}
