package product

type GetDto struct {
	Name           string `json:"name"`
	AvailableCount int    `json:"availableCount"`
	IsAvailable    bool   `json:"isAvailable"`
}
