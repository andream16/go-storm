package request

type Currency struct {
	Name string `json:"name"`
	Trend []CurrencyEntry
}

type CurrencyEntry struct {
	Date string `json:"date"`
	Value float64 `json:"value"`
}
