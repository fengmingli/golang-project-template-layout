package usecase

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/service"
)

type UpdateUserUseCase struct {
	userService *service.UserService
}

func NewUpdateUserUseCase(userService *service.UserService) *UpdateUserUseCase {
	return &UpdateUserUseCase{userService: userService}
}

func (u *UpdateUserUseCase) Execute(user *entity.User) error {
	return u.userService.UpdateUser(user)
}
