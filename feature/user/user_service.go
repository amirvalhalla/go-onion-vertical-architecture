package user

import (
	"context"

	user2 "github.com/amirvalhalla/go-onion-vertical-architecture/feature/user/dto"

	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/query"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetAll(ctx context.Context, pageIndex, pageSize int, search string) ([]domain.User, int64, error)
	Create(ctx context.Context, createDto user2.CreateDto) (domain.User, error)
	Update(ctx context.Context, updateDto user2.UpdateDto) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	uow sql.UnitOfWork
}

func NewService(uow sql.UnitOfWork) Service {
	return service{
		uow: uow,
	}
}

func (s service) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var userEntity domain.User
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if userEntity, err = uows.UserRepository().FindOne(query.FindByID(id)); err != nil {
			return err
		}
		return nil
	})

	return userEntity, err
}

func (s service) GetAll(ctx context.Context, pageIndex, pageSize int, search string) ([]domain.User, int64, error) {
	var userEntities []domain.User
	var totalRecords int64
	var err error

	err = s.uow.Do(ctx, false, func(uows sql.UnitOfWorkStore) error {
		if userEntities, totalRecords, err = uows.UserRepository().FindAllWithTotalRecords(
			query.WithUserSearch(search), // calc total records
			query.WithUserSearch(search),
			query.WithOrderBy("created_at", query.DESC),
			query.WithOffset(pageIndex, pageSize),
			query.WithLimit(pageSize),
		); err != nil {
			return err
		}

		return nil
	})

	return userEntities, totalRecords, err
}

func (s service) Create(ctx context.Context, createDto user2.CreateDto) (domain.User, error) {
	var err error
	userEntity := domain.NewUser(createDto.FirstName, createDto.LastName)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if userEntity, err = uows.UserRepository().Insert(userEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *userEntity, err
}

func (s service) Update(ctx context.Context, updateDto user2.UpdateDto) (domain.User, error) {
	var userEntity domain.User
	updatedUserEntity := new(domain.User)
	var err error

	if userEntity, err = s.Get(ctx, updateDto.ID); err != nil {
		return domain.User{}, err
	}

	userEntity.Update(updateDto.FirstName, updateDto.LastName, updateDto.Suspend)

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if updatedUserEntity, err = uows.UserRepository().Update(&userEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return *updatedUserEntity, err
}

func (s service) Delete(ctx context.Context, id uuid.UUID) error {
	var userEntity domain.User
	var err error

	if userEntity, err = s.Get(ctx, id); err != nil {
		return err
	}

	err = s.uow.Do(ctx, true, func(uows sql.UnitOfWorkStore) error {
		if err = uows.UserRepository().SoftDelete(&userEntity); err != nil {
			_ = uows.Rollback()
			return err
		}
		return uows.Commit()
	})

	return err
}
