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
	Create(userId int, comment common.Comment) (int, error)
	GetAll(userId int) ([]common.Comment, error)
	GetById(userId, commentId int) (common.Comment, error)
	Update(userId, commentId int, input common.UpdateCommentInput) error
	Delete(userId, commentId int) error
}

type Repository struct {
	Authorization
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Comment:       NewCommentPostgres(db),
	}
}
