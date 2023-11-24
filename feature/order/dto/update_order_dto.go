package order

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"github.com/google/uuid"
)

type UpdateDto struct {
	ID     uuid.UUID          `binding:"required" json:"id"`
	Status domain.OrderStatus `binding:"required" json:"status"`
}
