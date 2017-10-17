package request

import "time"

type Price struct {
	Item string `json:"item"`
	Price float64 `json:"price"`
	Date time.Time `json:"date"`
}

type Prices struct {
	Item string `json:"item"`
	Prices []PriceEntry `json:"prices"`
}

type PriceEntry struct {
	Price float64 `json:"price"`
	Date string `json:"date"`
}