package services

import (

	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	todo_collection *models.TodoCollection
}

func Todo(todo_collection *models.TodoCollection) *TodoService {
	return &TodoService{
		todo_collection: todo_collection,
	}
}

func (s TodoService) GetTodos()([]types.Todo, error){
	return s.todo_collection.GetTodos()
}

func (s TodoService) GetTodoByID(id primitive.ObjectID) (*types.Todo, error){
	return s.todo_collection.GetTodoByID(id)
}

func (s TodoService) CreateTodo(payload types.CreateTodoPayload) (*types.Todo, error) {
	todo := new(types.Todo)
	// {id:0,completed:false,body:""}
	todo.Body = payload.Body
	id, err := s.todo_collection.CreateTodo(*todo)
	if err != nil {return nil, err}

	todo.ID = id

	return todo, nil
}

// func EditTodo(payload types.CreateTodoPayload, id primitive.ObjectID) (*mongo.UpdateResult, error) {
// 	filter := bson.M{"_id":id}
// 	update := bson.M{"$set": bson.M{"body": payload.Body}}

// 	edt_res, err := models.Todos.UpdateOne(context.Background(), filter, update)
// 	if err != nil {return nil, err}

// 	return edt_res, nil
// }

// func ApproveTodo(id primitive.ObjectID) (error) {
// 	filter := bson.M{"_id":id}
// 	update := bson.M{"$set": bson.M{"approved": true}}

// 	_, err := models.Todos.UpdateOne(context.Background(), filter, update)
// 	if err != nil {return err}

// 	return nil
// }

// func CompleteTodo(id primitive.ObjectID) (error) {
// 	filter := bson.M{"_id":id}
// 	update := bson.M{"$set": bson.M{"completed": true}}

// 	_, err := models.Todos.UpdateOne(context.Background(), filter, update)
// 	if err != nil {return err}

// 	return nil
// }


// func DeleteTodo(id primitive.ObjectID) error {
// 	filter := bson.M{"_id":id}
// 	_, err := models.Todos.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }