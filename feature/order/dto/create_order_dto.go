package order

import "github.com/google/uuid"

type CreateDto struct {
	UserID     uuid.UUID   `binding:"required" json:"userId"`
	ProductIDs []uuid.UUID `binding:"required,gt=0" json:"ProductIds"`
}
