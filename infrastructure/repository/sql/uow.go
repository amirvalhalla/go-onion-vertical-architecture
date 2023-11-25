package sql

import (
	"context"
	"errors"

	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrSetupUnitOfWork = errors.New("could not setup unit of work")

type UnitOfWork interface {
	Do(ctx context.Context, isTransactional bool, fn UnitOfWorkBlock) error
	Setup(connString string) error
}

type unitOfWork struct {
	conn *gorm.DB
}

type UnitOfWorkStore interface {
	UserRepository() User
	OrderRepository() Order
	ProductRepository() Product
	Rollback() error
	Commit() error
}

type uowStore struct {
	userRepository    User
	orderRepository   Order
	productRepository Product
	tx                *gorm.DB
}

type UnitOfWorkBlock func(UnitOfWorkStore) error

func (u *uowStore) UserRepository() User {
	return u.userRepository
}

func (u *uowStore) OrderRepository() Order {
	return u.orderRepository
}

func (u *uowStore) ProductRepository() Product {
	return u.productRepository
}

func (u *uowStore) Rollback() error {
	return u.tx.Rollback().Error
}

func (u *uowStore) Commit() error {
	return u.tx.Commit().Error
}

func NewUnitOfWork() UnitOfWork {
	return &unitOfWork{}
}

func (uow *unitOfWork) Setup(connString string) error {
	var err error
	if uow.conn, err = gorm.Open(
		postgres.Open(connString),
		&gorm.Config{
			SkipDefaultTransaction:   true,
			PrepareStmt:              true,
			DisableNestedTransaction: true,
		},
	); err != nil {
		return ErrSetupUnitOfWork
	}

	return uow.autoMigrate()
}

func (uow *unitOfWork) Do(ctx context.Context, isTransactional bool, fn UnitOfWorkBlock) error {
	var tx *gorm.DB

	if isTransactional {
		tx = uow.conn.WithContext(ctx).Begin()
	} else {
		tx = uow.conn.WithContext(ctx)
	}

	uows := &uowStore{
		userRepository:    NewUserRepository(tx),
		orderRepository:   NewOrderRepository(tx),
		productRepository: NewProductRepository(tx),
		tx:                tx,
	}
	return fn(uows)
}

func (uow *unitOfWork) autoMigrate() error {
	return uow.conn.AutoMigrate(
		&domain.User{},
		&domain.Order{},
		&domain.Product{},
	)
}
