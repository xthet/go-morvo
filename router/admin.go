package router

import (
	"net/http"

	ctrls "github.com/xthet/go-morvo/controllers"
)

func AdminRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /admin/approve", ctrls.GetTodos)
	router.HandleFunc("DELETE /admin/todos", ctrls.GetTodos)
	router.HandleFunc("GET /admin/complete", ctrls.GetTodos)
}