package request

import "time"

type Currency struct {
	Name string `json:"name"`
	Trend []CurrencyEntry
}

type CurrencyEntry struct {
	Date time.Time `json:"date"`
	Value float64 `json:"value"`
}
