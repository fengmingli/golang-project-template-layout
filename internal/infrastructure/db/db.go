package db

import (
	"database/sql"

	"github.com/fengmingli/golang-project-template-layout/internal/infrastructure/config"
)

type Database struct {
	*sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}
