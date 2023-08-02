package todo

import (
	"github.com/theLemonday/k8s-todos/api/presenter"
	"github.com/theLemonday/k8s-todos/data/entities"
)

type Repository interface {
	CreateTodo(todo *entities.Todo) (*entities.Todo, error)
	GetTodoById(id string) (*presenter.Todo, error)
	GetAllTodos() (*[]presenter.Todo, error)
	UpdateTodo(todo *entities.Todo) (*entities.Todo, error)
	DeleteTodo(id string) error
}
