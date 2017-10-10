package request

import "time"

type Amazon struct {
	Item string `json:"item"`
	Manufacturer string `json:"manufacturer"`
	Review []Review
	Category []Category
}

type Review struct {
	Date time.Time `json:"date"`
	Text string `json:"text"`
	Sentiment uint `json:"sentiment"`
}

type Category struct {
	Name string `json:"category"`
}
