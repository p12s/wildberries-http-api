package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/wildberries-http-api"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (r *CommentPostgres) Create(userId int, comment common.Comment) (int, error) {
	tx, err := r.db.Begin() // по-хорошему здесь транзакция не нужна, таблица только одна. Но пусть будет, как пример
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id", commentsTable)
	row := tx.QueryRow(query, userId, comment.Title, comment.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CommentPostgres) GetAll(userId int) ([]common.Comment, error) {
	var comments []common.Comment

	query := fmt.Sprintf("SELECT user_id, title, description FROM %s WHERE user_id = $1", commentsTable)
	err := r.db.Select(&comments, query, userId)

	return comments, err
}

func (r *CommentPostgres) GetById(userId, commentId int) (common.Comment, error) {
	var comment common.Comment

	query := fmt.Sprintf(`SELECT id, user_id, title, description FROM %s WHERE user_id = $1 AND id = $2`, commentsTable)
	err := r.db.Get(&comment, query, userId, commentId)

	return comment, err
}

func (r *CommentPostgres) Update(userId, commentId int, input common.UpdateCommentInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, "title=$"+strconv.Itoa(argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, "description=$"+strconv.Itoa(argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE user_id = $%d AND id = $%d`,
		commentsTable, setQuery, argId, argId+1)
	args = append(args, userId, commentId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CommentPostgres) Delete(userId, commentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", commentsTable)
	_, err := r.db.Exec(query, userId, commentId)

	return err
}
