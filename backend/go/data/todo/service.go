package todo

import (
	"github.com/theLemonday/k8s-todos/api/presenter"
	"github.com/theLemonday/k8s-todos/data/entities"
)

type Service interface {
	CreateTodo(todo *entities.Todo) (*entities.Todo, error)
	GetTodoById(id string) (*presenter.Todo, error)
	GetAllTodos() (*[]presenter.Todo, error)
	UpdateTodo(todo *entities.Todo) (*entities.Todo, error)
	DeleteTodo(id string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTodo(todo *entities.Todo) (*entities.Todo, error) {
	return s.repo.CreateTodo(todo)
}

func (s *service) GetTodoById(id string) (*presenter.Todo, error) {
	return s.repo.GetTodoById(id)
}

func (s *service) GetAllTodos() (*[]presenter.Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *service) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	return s.repo.UpdateTodo(todo)
}

func (s *service) DeleteTodo(id string) error {
	return s.repo.DeleteTodo(id)
}
