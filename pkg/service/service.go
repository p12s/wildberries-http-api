package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
	"github.com/spf13/viper"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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
	GetById(id int) (common.User, error)
	Update(id int, input common.UpdateUserInput) error
	Delete(id int) error
}

type Service struct {
	Authorization
	Comment
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Comment:       NewCommentService(repos.Comment),
		User:          NewUserService(repos.User),
	}
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("db.salt"))))
}
