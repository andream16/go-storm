package request

type Item struct {
	Item string `json:"item"`
	Category string `json:"category,omitempty"`
	URL string `json:"url"`
	Image string `json:"image"`
	Title string `json:"title"`
	Description string `json:"description"`
}
