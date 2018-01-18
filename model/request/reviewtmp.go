package request

type ReviewTmpEntry struct {
	Item string `json:"item,omitempty"`
	Date string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	Sentiment float64 `json:"sentiment,omitempty"`
	Stars float64 `json:"stars,omitempty"`
}

type ReviewTmp struct {
	Item string `json:"item,omitempty"`
	Reviews []ReviewTmpEntry `json:"reviews"`
}