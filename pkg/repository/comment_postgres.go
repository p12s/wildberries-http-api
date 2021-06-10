package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/wildberries-http-api"
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

	query := fmt.Sprintf("SELECT user_id, title, description FROM %s WHERE id = $1 AND user_id = $2", commentsTable)
	err := r.db.Select(&comment, query, userId, commentId)

	return comment, err
}
