package common

import "errors"

type Comment struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type UpdateCommentInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u UpdateCommentInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structures hasn't values")
	}
	return nil
}
