package repository

import "github.com/fengmingli/golang-project-template-layout/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id int64) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) error
}
