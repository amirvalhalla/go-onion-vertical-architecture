package product

import (
	"context"

	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	dto "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/dto"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/query"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (domain.Product, error)
	GetProductsByID(ctx context.Context, productIds []uuid.UUID) ([]domain.Product, error)
	GetAll(ctx context.Context, pageIndex, pageSize int, search string) ([]domain.Product, int64, error)
	Create(ctx context.Context, createDto dto.CreateDto) (domain.Product, error)
	Update(ctx context.Context, updateDto dto.UpdateDto) (domain.Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	uow sql.UnitOfWork
}

func NewService(uow sql.UnitOfWork) Service {
	return &service{
		uow: uow,
	}
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (domain.Product, error) {
	var productEntity domain.Product
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if productEntity, err = uows.ProductRepository().FindOne(query.FindByID(id)); err != nil {
			return err
		}
		return nil
	})

	return productEntity, err
}

func (s *service) GetAll(ctx context.Context, pageIndex, pageSize int, search string) ([]domain.Product, int64, error) {
	var productEntities []domain.Product
	var totalRecords int64
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if productEntities, totalRecords, err = uows.ProductRepository().FindAllWithTotalRecords(
			query.WithProductSearch(search), // calc total records
			query.WithProductSearch(search),
			query.WithOrderBy("created_at", query.DESC),
			query.WithOffset(pageIndex, pageSize),
			query.WithLimit(pageSize),
		); err != nil {
			return err
		}

		return nil
	})

	return productEntities, totalRecords, err
}

func (s *service) GetProductsByID(ctx context.Context, productIds []uuid.UUID) ([]domain.Product, error) {
	var productEntities []domain.Product
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if productEntities, err = uows.ProductRepository().FindAll(
			query.FindProductsWithIDs(productIds),
			query.WithOrderBy("created_at", query.DESC),
		); err != nil {
			return err
		}

		return nil
	})

	return productEntities, err
}

func (s *service) Create(ctx context.Context, createDto dto.CreateDto) (domain.Product, error) {
	var err error
	productEntity := domain.NewProduct(createDto.Name, createDto.AvailableCount)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if productEntity, err = uows.ProductRepository().Insert(productEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *productEntity, err
}

func (s *service) Update(ctx context.Context, updateDto dto.UpdateDto) (domain.Product, error) {
	var productEntity domain.Product
	updatedProductEntity := new(domain.Product)
	var err error

	if productEntity, err = s.Get(ctx, updateDto.ID); err != nil {
		return domain.Product{}, err
	}

	productEntity.Update(updateDto.Name, updateDto.AvailableCount, updateDto.IsAvailable)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if updatedProductEntity, err = uows.ProductRepository().Update(&productEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *updatedProductEntity, err
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	var productEntity domain.Product
	var err error

	if productEntity, err = s.Get(ctx, id); err != nil {
		return err
	}

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if err = uows.ProductRepository().SoftDelete(&productEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return err
}
