package product

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	dto "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/dto"
)

func ToGetDto(p domain.Product) *dto.GetDto {
	return &dto.GetDto{
		Name:           p.Name,
		AvailableCount: p.AvailableCount,
		IsAvailable:    p.IsAvailable,
	}
}

func ToGetDtos(products []domain.Product) *[]dto.GetDto {
	productsDto := make([]dto.GetDto, 0, len(products))
	for _, p := range products {
		productsDto = append(productsDto, *ToGetDto(p))
	}
	return &productsDto
}
