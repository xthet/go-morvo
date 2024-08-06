package controllers

import (
	"fmt"
	"net/http"

	svcs "github.com/xthet/go-morvo/services"
	"github.com/xthet/go-morvo/types"
	"github.com/xthet/go-morvo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	all, err := svcs.GetTodos()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.Todo{"todos": all})
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	todo, err := svcs.GetTodo(object_id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]types.Todo{"todo": *todo})
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

func EditTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	var payload types.CreateTodoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Body == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("EMPTY BODY"))
		return
	}

	res, err := svcs.EditTodo(payload, object_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, *res)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	err = svcs.CompleteTodo(object_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message":"TODO completed succesffully"})
}

func ApproveTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	err = svcs.ApproveTodo(object_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message":"TODO approved succesffully"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	err = svcs.DeleteTodo(object_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message":"TODO deleted succesffully"})
}