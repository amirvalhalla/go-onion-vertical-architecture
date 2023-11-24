package order

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	product "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/dto"
	"github.com/google/uuid"
)

type GetDto struct {
	ID       uuid.UUID          `json:"id"`
	OrderID  int32              `json:"orderId"`
	Status   domain.OrderStatus `json:"status"`
	Products []product.GetDto   `json:"products"`
}
