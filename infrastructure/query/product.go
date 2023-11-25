package query

import (
	"strings"

	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func WithProductSearch(search string) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("LOWER(name) LIKE ?", util.StringToLikeQueryExpression(strings.ToLower(search)))
	}
}

func FindProductsWithIDs(productIds []uuid.UUID) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("product_id in (?)", productIds)
	}
}
