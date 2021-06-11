package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/wildberries-http-api"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user common.User) (int, error) {
	tx, err := r.db.Begin() // по-хорошему здесь транзакция не нужна, таблица только одна. Но пусть будет, как пример
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, username, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *UserPostgres) GetById(id int) (common.User, error) {
	var user common.User

	query := fmt.Sprintf(`SELECT id, name, email, username FROM %s WHERE id = $1`, usersTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) Update(id int, input common.UpdateUserInput) error {
	// добавление обновляемых строк сделано так для работы, если их будет несколько (description, time, tag, etc.)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, "name=$"+strconv.Itoa(argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, "email=$"+strconv.Itoa(argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, "username=$"+strconv.Itoa(argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, "password_hash=$"+strconv.Itoa(argId))
		args = append(args, *input.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`,
		usersTable, setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UserPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, id)

	return err
}
