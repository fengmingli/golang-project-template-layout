package service

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entity.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id int64) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(user *entity.User) error {
	return s.repo.Update(user)
}
