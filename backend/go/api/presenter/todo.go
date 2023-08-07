package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/theLemonday/k8s-todos/data/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID      primitive.ObjectID `json:"id"`
	Content string             `json:"content"`
}

type TodoRequest struct {
	Type MessageType `json:"type"`
	Todo
}

type TodoReponse struct {
	ErrorReponse `json:",omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func Entity2Presenter(todo *entities.Todo) *Todo {
	return &Todo{
		ID:      todo.ID,
		Content: todo.Content,
	}
}

func Presenter2Entity(todo *Todo) *entities.Todo {
	return &entities.Todo{
		ID:      todo.ID,
		Content: todo.Content,
	}
}

func (t TodoReponse) MarshalSuccessOne(todo *Todo) ([]byte, error) {
	return json.Marshal(&TodoReponse{
		Data: todo,
	})
}

func (t TodoReponse) MarshalSuccessMany(todos *[]Todo) ([]byte, error) {
	return json.Marshal(&TodoReponse{
		Data: todos,
	})
}

func (t TodoReponse) MarshalFailure(err *ErrorReponse) ([]byte, error) {
	return json.Marshal(&TodoReponse{
		ErrorReponse: *err,
	})
}

func (t TodoReponse) WriteToHttpSuccessOneTodo(w http.ResponseWriter, todo *Todo) {
	res, err := TodoReponse{}.MarshalSuccessOne(todo)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}

	w.Write(res)
}

func WriteReponseTodoToReponseWriter(w http.ResponseWriter, resTodo *Todo) {
	res, err := TodoReponse{}.MarshalSuccessOne(resTodo)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}

	w.Write(res)
}

func TodosSuccessReponse(w http.ResponseWriter, todos *[]Todo) {
	res, err := TodoReponse{}.MarshalSuccessMany(todos)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}

	w.Write(res)
}

func TodoErrorReponse(w http.ResponseWriter, err error) {
	res, err := json.Marshal(map[string]interface{}{
		"status": true,
		// "data": resTodos,
		"error": err.Error(),
	})
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}

	w.Write(res)
}
