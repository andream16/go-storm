package request

import "time"

type Trend struct {
	Manufacturer string `json:"manufacturer"`
	Trend []TrendEntry
}

type TrendEntry struct {
	Date time.Time `json:"date"`
	Value float64 `json:"value"`
}
