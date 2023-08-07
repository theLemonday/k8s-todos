package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/theLemonday/k8s-todos/api/presenter"
	"github.com/theLemonday/k8s-todos/data/entities"
	"github.com/theLemonday/k8s-todos/data/todo"
	usererror "github.com/theLemonday/k8s-todos/user_error"
)

type TodosResource struct {
	s todo.Service
}

func (rs TodosResource) Routes(s todo.Service) chi.Router {
	r := chi.NewRouter()
	rs.s = s

	r.Post("/", rs.Create)
	r.Get("/", rs.GetAllTodos)

	r.Route("/{todoId}", func(r chi.Router) {
		r.Get("/", rs.GetTodoById)
		r.Put("/", rs.UpdateTodo)
		r.Delete("/", rs.DeleteTodo)
	})

	return r
}

func (rs TodosResource) Create(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var todo entities.Todo

	err := dec.Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resTodo, err := rs.s.CreateTodo(&todo)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	presenter.SuccessTodoReponse(w, presenter.EntityToPresenter(resTodo))
}

func (rs TodosResource) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	resTodos, err := rs.s.GetAllTodos()
	if err != nil {
		msg := "Error get all todos"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Error().Msg(msg)
		return
	}

	presenter.TodosSuccessReponse(w, resTodos)
}

func (rs TodosResource) GetTodoById(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	resTodo, err := rs.s.GetTodoById(todoId)
	if err == usererror.ErrNoDocumentFound {
		msg := fmt.Sprintf("Todo with id %s not found", todoId)
		http.Error(w, msg, http.StatusNotFound)
		log.Info().Msg(msg)
		return
	}

	if err != nil {
		http.Error(w, "Server is busy", http.StatusInternalServerError)
		log.Warn().Err(err)
		return
	}

	presenter.SuccessTodoReponse(w, resTodo)
}

func (rs TodosResource) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var todo entities.Todo

	err := dec.Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resTodo, err := rs.s.UpdateTodo(&todo)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	presenter.SuccessTodoReponse(w, presenter.EntityToPresenter(resTodo))
}

func (rs TodosResource) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	err := rs.s.DeleteTodo(todoId)
	if err == usererror.ErrNoDocumentFound {
		msg := fmt.Sprintf("Todo with id %s not found", todoId)
		http.Error(w, msg, http.StatusNotFound)
		log.Info().Msg(msg)
		return
	}

	presenter.SuccessTodoReponse(w, nil)
}
