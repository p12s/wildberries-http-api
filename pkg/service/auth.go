package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
	"github.com/spf13/viper"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user common.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("db.salt"))))
}
