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

func (r *CommentPostgres) Create(idUser int, comment common.Comment) (int, error) {
	tx, err := r.db.Begin() // по-хорошему здесь транзакция не нужна, таблица только одна. Но пусть будет, как пример
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (id_user, txt) VALUES ($1, $2) RETURNING id", commentsTable)
	row := tx.QueryRow(query, idUser, comment.Txt)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CommentPostgres) GetAll(idUser int) ([]common.Comment, error) {
	var comments []common.Comment

	query := fmt.Sprintf("SELECT id, id_user, txt FROM %s WHERE id_user = $1", commentsTable)
	if err := r.db.Select(&comments, query, idUser); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentPostgres) GetById(idUser, commentId int) (common.Comment, error) {
	var comment common.Comment

	query := fmt.Sprintf(`SELECT id, id_user, txt FROM %s WHERE id_user = $1 AND id = $2`, commentsTable)
	if err := r.db.Get(&comment, query, idUser, commentId); err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *CommentPostgres) Update(idUser, commentId int, input common.UpdateCommentInput) error {
	// добавление обнволяемых строк сделано так для работы, если их будет несколько (description, time, tag, etc.)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Txt != nil {
		setValues = append(setValues, "txt=$"+strconv.Itoa(argId))
		args = append(args, *input.Txt)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id_user = $%d AND id = $%d`,
		commentsTable, setQuery, argId, argId+1)
	args = append(args, idUser, commentId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CommentPostgres) Delete(idUser, commentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_user = $1 AND id = $2", commentsTable)
	_, err := r.db.Exec(query, idUser, commentId)

	return err
}
