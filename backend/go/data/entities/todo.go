package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
