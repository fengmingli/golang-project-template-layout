package usecase

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/service"
)

type CreateUserUseCase struct {
	userService *service.UserService
}

func NewCreateUserUseCase(userService *service.UserService) *CreateUserUseCase {
	return &CreateUserUseCase{userService: userService}
}

func (u *CreateUserUseCase) Execute(user *entity.User) error {
	return u.userService.CreateUser(user)
}
