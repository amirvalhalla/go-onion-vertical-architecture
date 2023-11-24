package order

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	dto "github.com/amirvalhalla/go-onion-vertical-architecture/feature/order/dto"
	product "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/mapper"
)

func ToGetDto(p domain.Order) *dto.GetDto {
	return &dto.GetDto{
		ID:       p.ID,
		OrderID:  p.OrderID,
		Status:   p.Status,
		Products: *product.ToGetDtos(p.Products),
	}
}

func ToGetDtos(orders []domain.Order) *[]dto.GetDto {
	ordersDto := make([]dto.GetDto, 0, len(orders))
	for _, p := range orders {
		ordersDto = append(ordersDto, *ToGetDto(p))
	}
	return &ordersDto
}
