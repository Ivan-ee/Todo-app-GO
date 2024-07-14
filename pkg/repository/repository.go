package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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
