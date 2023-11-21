package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	Name           string    `gorm:"size:256"`
	AvailableCount int
	IsAvailable    bool      `gorm:"index"`
	OrderID        uuid.UUID `gorm:"index"`
	gorm.Model
}

func NewProduct(name string, count int, orderID uuid.UUID) *Product {
	return &Product{
		ID:             uuid.New(),
		Name:           name,
		AvailableCount: count,
		IsAvailable:    true,
		OrderID:        orderID,
	}
}

func (p *Product) ChangeName(name string) {
	p.Name = name
}

func (p *Product) ChangeAvailableCount(count int) {
	p.AvailableCount = count
}

func (p *Product) ChangeAvailability(status bool) {
	p.IsAvailable = status
}
