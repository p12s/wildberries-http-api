package common

import (
	"errors"
	"net/mail"
)

type Comment struct {
	Id     int    `json:"id" db:"id"`
	IdUser int    `json:"id_user" db:"id_user"`
	Txt    string `json:"txt" db:"txt" binding:"required"`
}

type UpdateCommentInput struct {
	Txt *string `json:"txt"`
}

func (u UpdateCommentInput) Validate() error {
	if u.Txt == nil {
		return errors.New("update structure hasn't values")
	}
	return nil
}

type UpdateUserInput struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

func (u UpdateUserInput) Validate() error {
	if u.Name == nil && u.Email == nil && u.Username == nil && u.Password == nil {
		return errors.New("update structure hasn't values")
	}
	if _, err := mail.ParseAddress(*u.Email); err != nil {
		return errors.New("update structure email isn't valid")
	}
	return nil
}
