package product

import "github.com/google/uuid"

type GetDto struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	AvailableCount int       `json:"availableCount"`
	IsAvailable    bool      `json:"isAvailable"`
}
