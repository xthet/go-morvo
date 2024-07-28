package services

import (
	"context"

	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos() ([]types.Todo, error){
	var todos []types.Todo

	cursor, err := models.Todos.Find(context.Background(), bson.M{})
	if err != nil {return nil, err}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo types.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func CreateTodo(payload types.CreateTodoPayload) (*types.Todo, error) {
	todo := new(types.Todo)
	// {id:0,completed:false,body:""}

	todo.Body = payload.Body

	ins_res, err := models.Todos.InsertOne(context.Background(), todo)
	if err != nil {return nil, err}

	todo.ID = ins_res.InsertedID.(primitive.ObjectID)

	return todo, nil
}