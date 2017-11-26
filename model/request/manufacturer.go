package request

type Manufacturer struct{
	Name string `json:"name,omitempty"`
}

type ManufacturerRequest struct {
	Manufacturer string `json:"manufacturer,omitempty"`
}

type Manufacturers struct {
	Manufacturers []Manufacturer `json:"manufacturers"`
}