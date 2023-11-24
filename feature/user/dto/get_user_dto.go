package user

import "github.com/google/uuid"

type GetUserDto struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	// TODO should implement orders dto here
}
