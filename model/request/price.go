package request

import "time"

type Price struct {
	Item string `json:"item"`
	Prices []Price
}

type PriceEntry struct {
	Price float64 `json:"price"`
	Date time.Time `json:"date"`
}