package order

import (
	"context"
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	dto "github.com/amirvalhalla/go-onion-vertical-architecture/feature/order/dto"
	"github.com/amirvalhalla/go-onion-vertical-architecture/feature/product"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/query"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (domain.Order, error)
	GetAll(ctx context.Context, pageIndex, pageSize int, search int) ([]domain.Order, int64, error)
	Create(ctx context.Context, createDto dto.CreateDto) (domain.Order, error)
	Update(ctx context.Context, updateDto dto.UpdateDto) (domain.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	uow            sql.UnitOfWork
	productService product.Service
}

func NewService(uow sql.UnitOfWork) Service {
	return &service{
		uow:            uow,
		productService: product.NewService(uow),
	}
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	var orderEntity domain.Order
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if orderEntity, err = uows.OrderRepository().FindOne(query.FindByID(id)); err != nil {
			return err
		}
		return nil
	})

	return orderEntity, err
}

func (s *service) GetAll(ctx context.Context, pageIndex, pageSize int, search int) ([]domain.Order, int64, error) {
	var orderEntities []domain.Order
	var totalRecords int64
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if orderEntities, totalRecords, err = uows.OrderRepository().FindAllWithTotalRecords(
			query.WithOrderSearch(search), // calc total records
			query.WithOrderSearch(search),
			query.WithOrderBy("created_at", query.DESC),
			query.WithOffset(pageIndex, pageSize),
			query.WithLimit(pageSize),
		); err != nil {
			return err
		}

		return nil
	})

	return orderEntities, totalRecords, err
}

func (s *service) Create(ctx context.Context, createDto dto.CreateDto) (domain.Order, error) {
	var err error

	products, err := s.productService.GetProductsByID(ctx, createDto.ProductIDs)
	if err != nil {
		return domain.Order{}, err
	}

	orderEntity := domain.NewOrder(createDto.UserID, products)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if orderEntity, err = uows.OrderRepository().Insert(orderEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *orderEntity, err
}

func (s *service) Update(ctx context.Context, updateDto dto.UpdateDto) (domain.Order, error) {
	var orderEntity domain.Order
	updatedOrderEntity := new(domain.Order)
	var err error

	if orderEntity, err = s.Get(ctx, updateDto.ID); err != nil {
		return domain.Order{}, err
	}

	orderEntity.ChangeStatus(updateDto.Status)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if updatedOrderEntity, err = uows.OrderRepository().Update(&orderEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *updatedOrderEntity, err
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	var orderEntity domain.Order
	var err error

	if orderEntity, err = s.Get(ctx, id); err != nil {
		return err
	}

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if err = uows.OrderRepository().SoftDelete(&orderEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return err
}
