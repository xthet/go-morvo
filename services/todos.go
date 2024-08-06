package services

import (
	"context"

	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetTodo(id primitive.ObjectID) (*types.Todo, error){
	filter := bson.M{"_id":id}
	todo := new(types.Todo)
	err := models.Todos.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {return nil, err}
	return todo, nil
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

func EditTodo(payload types.CreateTodoPayload, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"body": payload.Body}}

	edt_res, err := models.Todos.UpdateOne(context.Background(), filter, update)
	if err != nil {return nil, err}

	return edt_res, nil
}

func ApproveTodo(id primitive.ObjectID) (error) {
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"approved": true}}

	_, err := models.Todos.UpdateOne(context.Background(), filter, update)
	if err != nil {return err}

	return nil
}

func CompleteTodo(id primitive.ObjectID) (error) {
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err := models.Todos.UpdateOne(context.Background(), filter, update)
	if err != nil {return err}

	return nil
}


func DeleteTodo(id primitive.ObjectID) error {
	filter := bson.M{"_id":id}
	_, err := models.Todos.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}