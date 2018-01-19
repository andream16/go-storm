package request

type ReviewEntry struct {
	Item string `json:"item,omitempty"`
	Date string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	Sentiment float64 `json:"sentiment,omitempty"`
	Stars float64 `json:"stars,omitempty"`
}

type Review struct {
	Item string `json:"item,omitempty"`
	Reviews []ReviewEntry `json:"reviews"`
}