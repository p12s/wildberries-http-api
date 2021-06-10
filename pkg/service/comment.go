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

func (s *CommentService) Create(idUser int, comment common.Comment) (int, error) {
	return s.repo.Create(idUser, comment)
}

func (s *CommentService) GetAll(idUser int) ([]common.Comment, error) {
	return s.repo.GetAll(idUser)
}

func (s *CommentService) GetById(idUser, commentId int) (common.Comment, error) {
	return s.repo.GetById(idUser, commentId)
}

func (s *CommentService) Update(idUser, commentId int, input common.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(idUser, commentId, input)
}

func (s *CommentService) Delete(idUser, listId int) error {
	return s.repo.Delete(idUser, listId)
}
