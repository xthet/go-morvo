package models

import (
	"context"

	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoCollection struct {
	collection *mongo.Collection
}

func Todo(collection *mongo.Collection) *TodoCollection {
	return &TodoCollection{
		collection: collection,
	}
}

func (c TodoCollection) CreateTodo(todo types.Todo) (primitive.ObjectID, error) {
	ins_res, err := c.collection.InsertOne(context.Background(), todo)
	ID := ins_res.InsertedID.(primitive.ObjectID)
	return ID, err
}

func (c TodoCollection) GetTodoByID(id primitive.ObjectID) (*types.Todo, error) {
	filter := bson.M{"_id":id}
	todo := new(types.Todo)
	
	// if err, then err is passed to err, if not, the found todo is used to modify todo
	err := c.collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {return nil, err}
	return todo, nil
}

func (c TodoCollection) GetTodos()([]types.Todo, error){
	var todos []types.Todo

	cursor, err := c.collection.Find(context.Background(), bson.M{})
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