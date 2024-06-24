package usecase

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/service"
)

type GetUserUseCase struct {
	userService *service.UserService
}

func NewGetUserUseCase(userService *service.UserService) *GetUserUseCase {
	return &GetUserUseCase{userService: userService}
}

func (u *GetUserUseCase) Execute(id int64) (*entity.User, error) {
	return u.userService.GetUser(id)
}
