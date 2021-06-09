package service

import (
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "s11212-df2309dsfso1[9wasdf1[0wf23"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
