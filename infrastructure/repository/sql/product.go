package sql

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"gorm.io/gorm"
)

type Product interface {
	BaseRepository[domain.Product]
}

func NewProductRepository(tx *gorm.DB) Product {
	return NewBaseRepository[domain.Product](tx)
}
