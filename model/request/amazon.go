package request


type Amazon struct {
	Item Item `json:"item"`
	Manufacturer ManufacturerRequest `json:"manufacturer"`
	Review Review `json:"review"`
	Category CategoryRequest `json:"categories"`
}
