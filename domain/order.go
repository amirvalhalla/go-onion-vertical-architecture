package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus int8

const (
	OrderPending OrderStatus = iota
	OrderProcessing
	OrderSending
	OrderReceived
	OrderCanceled
)

type Order struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	OrderID  int32     `gorm:"autoIncrement"`
	Status   OrderStatus
	UserID   uuid.UUID `gorm:"index"`
	Products []Product `gorm:"many2many:order_products"`
	gorm.Model
}

func NewOrder(userID uuid.UUID, products []Product) *Order {
	return &Order{
		ID:       uuid.New(),
		Status:   OrderPending,
		UserID:   userID,
		Products: products,
	}
}

func (o *Order) ChangeStatus(status OrderStatus) {
	o.Status = status
}
