package router

import (
	"fmt"
	"net/http"

	ctrls "github.com/xthet/go-morvo/controllers"
	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/services"
	"go.mongodb.org/mongo-driver/mongo"
)


func TodoRoutes(router *http.ServeMux, client *mongo.Client){
	todo_controller := ctrls.Todo(services.Todo(models.Todo(client.	Database("morvo").Collection("todos"))))
	fmt.Println("created again")

	// TO-DO ROUTES
	router.HandleFunc("GET /todos", todo_controller.GetTodos)
	router.HandleFunc("GET /todos/{id}", todo_controller.GetTodoByID)
	// router.HandleFunc("GET /todos/complete", ctrls.CompleteTodo)
	router.HandleFunc("POST /todos", todo_controller.CreateTodo)
	// router.HandleFunc("PATCH /todos", ctrls.EditTodo)
	// router.HandleFunc("DELETE /todos", ctrls.DeleteTodo)
}