package query

import (
	"fmt"

	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindByID(id uuid.UUID) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("id = ?", id)
	}
}

func FindByCustomColumn[T any](columnName string, value T) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where(fmt.Sprintf("%s = ?", columnName), value)
	}
}

func WithOffset(pageIndex, pageSize int) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Offset(util.CalcOffset(pageIndex, pageSize))
	}
}

func WithLimit(pageSize int) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Limit(pageSize)
	}
}

func WithOrderBy(columnName string, sort SortType) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Order(columnName + " " + sort.String())
	}
}

func WithoutAnySearch() sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q
	}
}

func WithNullColumnCondition(columnName string) sql.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where(fmt.Sprintf("%s IS NULL", columnName)).
			Or(fmt.Sprintf("%s = ?", columnName), "")
	}
}
