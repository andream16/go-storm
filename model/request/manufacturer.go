package request

type Manufacturer struct{
	Name string `json:"name,omitempty"`
}

type ManufacturerRequest struct {
	Item string `json:"item,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
}
