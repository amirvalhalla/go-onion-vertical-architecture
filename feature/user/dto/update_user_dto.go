package user

import "github.com/google/uuid"

type UpdateUserDto struct {
	ID        uuid.UUID `binding:"required" json:"id"`
	FirstName string    `binding:"required,min=3,max=128" json:"firstName"`
	LastName  string    `binding:"required,min=3,max=128" json:"lastName"`
	Suspend   bool      `binding:"required" json:"suspend"`
}
