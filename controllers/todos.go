package controllers

import (
	"fmt"
	"net/http"

	svcs "github.com/xthet/go-morvo/services"
	"github.com/xthet/go-morvo/types"
	"github.com/xthet/go-morvo/utils"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	all, err := svcs.GetTodos()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.Todo{"todos": all})
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateTodoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Body == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("EMPTY BODY"))
		return
	}

	res, err := svcs.CreateTodo(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, *res)
}
