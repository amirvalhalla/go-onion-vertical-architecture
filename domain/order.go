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
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	OrderID int32     `gorm:"autoIncrement"`
	Status  OrderStatus
	UserID  uuid.UUID `gorm:"index"`
	Product Product
	gorm.Model
}

func NewOrder(userID uuid.UUID) *Order {
	return &Order{
		ID:     uuid.New(),
		Status: OrderPending,
		UserID: userID,
	}
}

func (o *Order) ChangeStatus(status OrderStatus) {
	o.Status = status
}
