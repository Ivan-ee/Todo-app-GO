package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Update(userId int, listId int, input todo.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface{}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: newAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
