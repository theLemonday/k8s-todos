package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Content   string             `json:"content"`
	CreatedAt time.Time          `json:"created_at"`
	UpdateAt  time.Time          `json:"updated_at"`
}
