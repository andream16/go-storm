package request


type Amazon struct {
	Item string `json:"item,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Review []Review
	Category []AmazonCategory
}

type ReviewEntry struct {
	Date string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	Sentiment uint `json:"sentiment,omitempty"`
	Stars uint `json:"stars,omitempty"`
}

type Review struct {
	Item string `json:"item,omitempty"`
	Review []ReviewEntry `json:"review"`
}

type AmazonCategory struct {
	Name string `json:"category,omitempty"`
}
