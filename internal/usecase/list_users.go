package usecase

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/service"
)

type ListUsersUseCase struct {
	userService *service.UserService
}

func NewListUsersUseCase(userService *service.UserService) *ListUsersUseCase {
	return &ListUsersUseCase{userService: userService}
}

func (u *ListUsersUseCase) Execute() ([]*entity.User, error) {
	return u.userService.ListUsers()
}
