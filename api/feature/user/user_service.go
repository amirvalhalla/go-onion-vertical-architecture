package user

import (
	"context"
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
}

func NewService() Service {
	return service{}
}

func (s service) Get(ctx context.Context, id uuid.UUID) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetAll(ctx context.Context, pageIndex, pageSize int, search string) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(ctx context.Context, createDto user.CreateUserDto) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
