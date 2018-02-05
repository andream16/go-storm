package request

type Statistics struct {
	Item string `json:"item,omitempty"`
	Name string   `json:"name,omitempty"`
	Score float64 `json:"score,omitempty"`
	TestSize string `json:"test_size,omitempty"`
	Forecast []StatisticsEntry `json:"forecast_entries"`
}

type StatisticsEntry struct {
	Price float64 `json:"price,omitempty"`
	Date string`json:"date,omitempty"`
	Score float64 `json:"score,omitempty"`
	TestSize string `json:"test_size,omitempty"`
}
