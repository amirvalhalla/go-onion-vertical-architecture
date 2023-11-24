package user

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/domain"
	dto "github.com/amirvalhalla/go-onion-vertical-architecture/feature/user/dto"
)

func ToGetDto(u domain.User) *dto.GetUserDto {
	return &dto.GetUserDto{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func ToGetDtos(users []domain.User) *[]dto.GetUserDto {
	usersDto := make([]dto.GetUserDto, 0, len(users))
	for _, u := range users {
		usersDto = append(usersDto, *ToGetDto(u))
	}
	return &usersDto
}
