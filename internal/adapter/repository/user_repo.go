package repository

import (
	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/repository"
	"github.com/fengmingli/golang-project-template-layout/internal/infrastructure/db"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (r *UserRepository) FindByID(id int64) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	return nil
}
