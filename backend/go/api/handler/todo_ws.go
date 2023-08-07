package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/theLemonday/k8s-todos/api/presenter"
	"github.com/theLemonday/k8s-todos/data/todo"
)

type TodoWebsocketHandler struct {
	s todo.Service
}

func NewTodoWebsocketHandler(s todo.Service) *TodoWebsocketHandler {
	return &TodoWebsocketHandler{
		s: s,
	}
}

func (h *TodoWebsocketHandler) HandleMessage(msg []byte) []byte {
	var todoMsg presenter.TodoRequest
	err := json.Unmarshal(msg, &todoMsg)
	if err != nil {
		log.Error().Err(err)
		// TODO unknown for now
		return []byte{}
	}

	var res []byte
	var resErr *presenter.ErrorReponse
	switch todoMsg.Type {
	case presenter.Create:
		res, resErr = h.create(&todoMsg.Todo)
	case presenter.GetAll:
		h.getAllTodos()
	default:
		return nil
	}

	if resErr != nil {
		return presenter.TodoReponse{}.MarshalFailure(resErr)
	}
	return nil
}

func (h *TodoWebsocketHandler) create(todo *presenter.Todo) ([]byte, *presenter.ErrorReponse) {
	Etodo, err := h.s.CreateTodo(presenter.Presenter2Entity(todo))
	if err != nil {
		return nil, &presenter.ErrorReponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	res, err := presenter.TodoReponse{}.MarshalSuccessOne(presenter.Entity2Presenter(Etodo))
	if err != nil {
		return nil, &presenter.ErrorReponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	return res, nil
}

func (h *TodoWebsocketHandler) getAllTodos() []byte {
	todos, err := h.s.GetAllTodos()
	if err != nil {
		presenter.TodoReponse{}.MarshalFailure(presenter.ErrorReponse{
			StatusCode: http.StatusInternalServerError,
		})
	}

	res, err := presenter.TodoReponse{}.MarshalSuccessMany(todos)

	return res
}
