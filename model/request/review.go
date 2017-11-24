package request

type ReviewEntry struct {
	Item string `json:"item,omitempty"`
	Date string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	Sentiment uint `json:"sentiment,omitempty"`
	Stars uint `json:"stars,omitempty"`
}

type Review struct {
	Item string `json:"item,omitempty"`
	Reviews []ReviewEntry `json:"reviews"`
}