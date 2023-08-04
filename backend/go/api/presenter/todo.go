package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/theLemonday/k8s-todos/data/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Content string             `json:"content"`
}

func EntityToPresenter(todo *entities.Todo) *Todo {
	return &Todo{
		ID:      todo.ID,
		Content: todo.Content,
	}
}

func SuccessTodoReponse(w http.ResponseWriter, resTodo *Todo) {
	res, err := json.Marshal(map[string]interface{}{
		"status": true,
		"data":   resTodo,
		// "error": "",
	})
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}

	w.Write(res)
}

func TodosSuccessReponse(w http.ResponseWriter, resTodos *[]Todo) {
	res, err := json.Marshal(map[string]interface{}{
		"status": true,
		"data":   resTodos,
		// "error": "",
	})
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
