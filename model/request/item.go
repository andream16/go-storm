package request

type Item struct {
	Item string `json:"item,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	HasReviews bool `json:"has_reviews,omitempty"`
	URL string `json:"url,omitempty"`
	Image string `json:"image,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
