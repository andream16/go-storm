package request

import "time"

type Amazon struct {
	Item string `json:"item"`
	Manufacturer string `json:"manufacturer"`
	Review []Review
	Category []AmazonCategory
}

type Review struct {
	Date time.Time `json:"date"`
	Text string `json:"text"`
	Sentiment uint `json:"sentiment"`
}

type AmazonCategory struct {
	Name string `json:"category"`
}
