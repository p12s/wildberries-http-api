package service

import (
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
}

type Comment interface {
}

type Service struct {
	Authorization
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
