package product

import "github.com/google/uuid"

type UpdateDto struct {
	ID             uuid.UUID `binding:"required" json:"id"`
	Name           string    `binding:"required,min=3,max=128" json:"name"`
	AvailableCount int       `binding:"required,gte=1" json:"availableCount"`
	IsAvailable    bool      `binding:"required" json:"isAvailable"`
}
