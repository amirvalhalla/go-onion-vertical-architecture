package sql

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"gorm.io/gorm"
)

type User interface {
	BaseRepository[domain.User]
}

func NewUserRepository(tx *gorm.DB) User {
	return NewBaseRepository[domain.User](tx)
}
