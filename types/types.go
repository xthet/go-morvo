package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Approved  bool               `json:"approved"`
	Body      string             `json:"body"`
}

type CreateTodoPayload struct {
	Body string `json:"body"`
}

type UpdateTodoPayload struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Password  string             `json:"-"`
	CreatedAt time.Time          `json:"createdAt"`
}

type RegisterUserResponse struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}
