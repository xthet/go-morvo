package router

import (
	"net/http"
	ctrls "github.com/xthet/go-morvo/controllers"
)


func TodoRoutes(router *http.ServeMux){
	// TO-DO ROUTES
	router.HandleFunc("GET /todos", ctrls.GetTodos)
	router.HandleFunc("GET /todos/{id}", ctrls.GetTodo)
	router.HandleFunc("GET /todos/complete", ctrls.CompleteTodo)
	router.HandleFunc("POST /todos", ctrls.CreateTodo)
	router.HandleFunc("PATCH /todos", ctrls.EditTodo)
	router.HandleFunc("DELETE /todos", ctrls.DeleteTodo)
}