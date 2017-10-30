package request


type Trend struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Trend []TrendEntry
}

type TrendEntry struct {
	Date string `json:"date,omitempty"`
	Value float64 `json:"value,omitempty"`
}
