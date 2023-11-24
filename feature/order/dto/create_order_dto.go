package order

import "github.com/google/uuid"

type CreateDto struct {
	UserID uuid.UUID `binding:"required" json:"userId"`
}
