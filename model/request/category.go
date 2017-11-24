package request

type Category struct {
	Category string `json:"category,omitempty"`
	Item string `json:"item,omitempty"`
}

type CategoryRequest struct {
	Item string `json:"item,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}
