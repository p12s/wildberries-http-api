package service

import (
	common "github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(userId int, comment common.Comment) (int, error) {
	return s.repo.Create(userId, comment)
}

func (s *CommentService) GetAll(userId int) ([]common.Comment, error) {
	return s.repo.GetAll(userId)
}

func (s *CommentService) GetById(userId, listId int) (common.Comment, error) {
	return s.repo.GetById(userId, listId)
}
