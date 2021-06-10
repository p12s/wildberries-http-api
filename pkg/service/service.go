package service

import (
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Comment interface {
	Create(userId int, comment common.Comment) (int, error)
	GetAll(userId int) ([]common.Comment, error)
	GetById(userId, listId int) (common.Comment, error)
}

type Service struct {
	Authorization
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Comment:       NewCommentService(repos.Comment),
	}
}
