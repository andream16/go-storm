package request


type Amazon struct {
	Item Item `json:"item"`
	Manufacturer Manufacturer `json:"manufacturer"`
	Review Review `json:"review"`
	Category CategoryRequest `json:"category"`
}
