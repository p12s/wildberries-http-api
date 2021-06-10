package service

import (
	"fmt"
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

func (s *CommentService) GetById(userId, commentId int) (common.Comment, error) {
	fmt.Println("service/comment.go userId commentId", userId, commentId)
	return s.repo.GetById(userId, commentId)
}

func (s *CommentService) Update(userId, commentId int, input common.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, commentId, input)
}

func (s *CommentService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
