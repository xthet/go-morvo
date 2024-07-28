package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool `json:"completed"`
	Approved bool `json:"approved"`
	Body string `json:"body"`
}

type CreateTodoPayload struct {
	Body string `json:"body"`
}