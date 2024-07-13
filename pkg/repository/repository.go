package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: newAuthPostgres(db),
	}
}
