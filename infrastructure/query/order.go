package query

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"gorm.io/gorm"
)

func WithOrderSearch(search int) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("order_status = ?", search)
	}
}
