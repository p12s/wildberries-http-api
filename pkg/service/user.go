package service

import (
	"fmt"
	common "github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user common.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.Create(user)
}

func (s *UserService) GetById(id int) (common.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(id int, input common.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	fmt.Println(*input.Password)
	*input.Password = GeneratePasswordHash(*input.Password)
	fmt.Println(*input.Password)
	return s.repo.Update(id, input)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
