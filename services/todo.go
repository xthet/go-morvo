package services

import (
	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s TodoService) EditTodo(payload types.CreateTodoPayload, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	return s.todo_collection.EditTodo(payload, id)
}

func (s TodoService) ApproveTodo(id primitive.ObjectID) (error) {
	return s.todo_collection.ApproveTodo(id)
}

func (s TodoService) CompleteTodo(id primitive.ObjectID) (error) {
	return s.todo_collection.CompleteTodo(id)
}


func (s TodoService) DeleteTodo(id primitive.ObjectID) error {
	return s.todo_collection.DeleteTodo(id)
}