package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	FirstName string    `gorm:"size:128"`
	LastName  string    `gorm:"size:128"`
	Orders    []Order
	Suspended bool
	gorm.Model
}

func NewUser(firstName, lastName string) *User {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Suspended: false,
	}
}

func (u *User) Update(firstName, lastName string, suspend bool) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Suspended = suspend
}

func (u *User) IsSuspended() bool {
	return u.Suspended
}

func (u *User) Suspend() {
	u.Suspended = true
}
