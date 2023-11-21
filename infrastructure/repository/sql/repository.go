package sql

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrFindOne             = errors.New("could not find one because of database error")
	ErrFindAll             = errors.New("could not find all because of database error")
	ErrRecordNotFound      = errors.New("could not found the record in database")
	ErrCouldNotInsert      = errors.New("could not insert model into database")
	ErrCouldNotBatchInsert = errors.New("could not batch insert models into database")
	ErrCouldNotUpdate      = errors.New("could not update model into database")
	ErrCouldNotSave        = errors.New("could not save model into database")
	ErrCouldNotDelete      = errors.New("could not delete model in database")
)

type SelectQuery func(*gorm.DB) *gorm.DB

type BaseRepository[T any] interface {
	FindOne(...SelectQuery) (T, error)
	FindAll(...SelectQuery) ([]T, error)
	FindAllWithTotalRecords(countQuery SelectQuery, sc ...SelectQuery) ([]T, int64, error)
	Insert(*T) (*T, error)
	BatchInsert([]*T) ([]*T, error)
	Update(*T) (*T, error)
	Save(*T) (*T, error)
	Delete(*T) error
	SoftDelete(*T) error
}

type baseRepository[T any] struct {
	tx *gorm.DB
}

func NewBaseRepository[T any](tx *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		tx: tx,
	}
}

func (b baseRepository[T]) FindOne(sc ...SelectQuery) (T, error) {
	var row T

	for i := range sc {
		b.tx = sc[i](b.tx)
	}

	if err := b.tx.First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return row, ErrRecordNotFound
		}
		return row, ErrFindOne
	}

	return row, nil
}

func (b baseRepository[T]) FindAll(sc ...SelectQuery) ([]T, error) {
	var rows []T

	for i := range sc {
		b.tx = sc[i](b.tx)
	}

	if err := b.tx.Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrFindAll
	}

	return rows, nil
}

func (b baseRepository[T]) FindAllWithTotalRecords(countQuery SelectQuery, sc ...SelectQuery) ([]T, int64, error) {
	var t T
	var rows []T
	var totalRecords int64

	countQuery(b.tx).Model(&t).Count(&totalRecords)

	for i := range sc {
		b.tx = sc[i](b.tx)
	}

	if err := b.tx.Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, ErrRecordNotFound
		}
		return nil, 0, ErrFindAll
	}

	return rows, totalRecords, nil
}

func (b baseRepository[T]) Insert(model *T) (*T, error) {
	if err := b.tx.Create(model).Error; err != nil {
		return nil, ErrCouldNotInsert
	}
	return model, nil
}

func (b baseRepository[T]) BatchInsert(models []*T) ([]*T, error) {
	if err := b.tx.Create(models).Error; err != nil {
		return nil, ErrCouldNotBatchInsert
	}
	return models, nil
}

func (b baseRepository[T]) Update(model *T) (*T, error) {
	if err := b.tx.Updates(model).Error; err != nil {
		return nil, ErrCouldNotUpdate
	}

	return model, nil
}

func (b baseRepository[T]) Save(model *T) (*T, error) {
	if err := b.tx.Save(model).Error; err != nil {
		return nil, ErrCouldNotSave
	}
	return model, nil
}

func (b baseRepository[T]) Delete(model *T) error {
	if err := b.tx.Unscoped().Delete(model).Error; err != nil {
		return ErrCouldNotDelete
	}
	return nil
}

func (b baseRepository[T]) SoftDelete(model *T) error {
	if err := b.tx.Delete(model).Error; err != nil {
		return ErrCouldNotDelete
	}
	return nil
}
