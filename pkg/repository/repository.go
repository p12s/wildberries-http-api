package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/wildberries-http-api"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GetUser(username, password string) (common.User, error)
}

type Comment interface {
	Create(idUser int, comment common.Comment) (int, error)
	GetAll(idUser int) ([]common.Comment, error)
	GetById(idUser, commentId int) (common.Comment, error)
	Update(idUser, commentId int, input common.UpdateCommentInput) error
	Delete(idUser, commentId int) error
}

type User interface {
	Create(user common.User) (int, error)
	GetById(idUser int) (common.User, error)
	Update(id int, input common.UpdateUserInput) error
	Delete(idUser int) error
}

type Repository struct {
	Authorization
	Comment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Comment:       NewCommentPostgres(db),
		User:          NewUserPostgres(db),
	}
}
