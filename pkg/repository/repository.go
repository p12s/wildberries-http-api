package repository

import (
	"github.com/p12s/wildberries-http-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user User) (int, error)
}

type Comment interface {
}

type Repository struct {
	Authorization
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
