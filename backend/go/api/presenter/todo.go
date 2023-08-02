package presenter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Content string             `json:"content"`
}
