package user

import (
	"context"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"

	user "github.com/amirvalhalla/go-onion-vertical-architecture/api/feature/user/dto"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID)
	GetAll(ctx context.Context, pageIndex, pageSize int, search string)
	Create(ctx context.Context, createDto user.CreateUserDto)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type service struct {
	uow sql.UnitOfWork
}

func NewService(uow sql.UnitOfWork) Service {
	return service{
		uow: uow,
	}
}

func (s service) Get(ctx context.Context, id uuid.UUID) {
	// TODO implement me
	panic("implement me")
}

func (s service) GetAll(ctx context.Context, pageIndex, pageSize int, search string) {
	// TODO implement me
	panic("implement me")
}

func (s service) Create(ctx context.Context, createDto user.CreateUserDto) {
	// TODO implement me
	panic("implement me")
}

func (s service) Update(ctx context.Context) {
	// TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx context.Context) {
	// TODO implement me
	panic("implement me")
}
