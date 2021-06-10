package common

import "errors"

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
