package service

import (
	"github.com/doffy007/banking-api/pkg/core/v1/user/dto"
	"github.com/doffy007/banking-api/pkg/core/v1/user/resource"
)

type InterfaceUserService interface {
	CreateUser(payload *dto.CreateUserDTO) error
	GetUser(payload *dto.GetUserDTO) (*dto.UserWithTokenDTO, error)
}

type UserService struct {
	UserResource resource.InterfaceUserResource
}

func InitUserService(userResource resource.InterfaceUserResource) InterfaceUserService {
	return &UserService{
		UserResource: userResource,
	}
}
