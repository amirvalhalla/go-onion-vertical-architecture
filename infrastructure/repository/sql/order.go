package sql

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"gorm.io/gorm"
)

type Order interface {
	BaseRepository[domain.Order]
}

func NewOrderRepository(tx *gorm.DB) Order {
	return NewBaseRepository[domain.Order](tx)
}
