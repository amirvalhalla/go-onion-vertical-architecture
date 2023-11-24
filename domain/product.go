package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	Name           string    `gorm:"size:256"`
	AvailableCount int
	IsAvailable    bool `gorm:"index"`
	gorm.Model
}

func NewProduct(name string, count int) *Product {
	return &Product{
		ID:             uuid.New(),
		Name:           name,
		AvailableCount: count,
		IsAvailable:    true,
	}
}

func (p *Product) Update(name string, availableCount int, isAvailable bool) {
	p.ChangeName(name)
	p.ChangeAvailableCount(availableCount)
	p.ChangeAvailability(isAvailable)
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
