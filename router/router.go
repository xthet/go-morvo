package router

import (
	"net/http"

	ctrls "github.com/xthet/go-morvo/controllers"
	"github.com/xthet/go-morvo/utils"
)

func Routes() http.Handler {
	router := http.NewServeMux()


	router.HandleFunc("GET /", greet)


	// TO-DO ROUTES
	router.HandleFunc("GET /todos", ctrls.GetTodos)
	router.HandleFunc("POST /todos", ctrls.CreateTodo)
	router.HandleFunc("DELETE /todos", ctrls.DeleteTodo)

	sub_router := http.NewServeMux()
	sub_router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))	

	return sub_router
}

func greet(w http.ResponseWriter, r *http.Request){
	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg":"Hello World"})
}
