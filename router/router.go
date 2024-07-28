package router

import (
	"net/http"

	"github.com/gorilla/mux"
	ctrls "github.com/xthet/go-morvo/controllers"
	"github.com/xthet/go-morvo/utils"
)

func Routes() http.Handler {
	grand_router := mux.NewRouter()
	router := grand_router.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/", greet).Methods("GET")


	// TO-DO ROUTES
	router.HandleFunc("/todos", ctrls.GetTodos).Methods("GET")
	router.HandleFunc("/todos", ctrls.CreateTodo).Methods("POST")

	return router
}

func greet(w http.ResponseWriter, r *http.Request){
	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg":"Hello World"})
}
