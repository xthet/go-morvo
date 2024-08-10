package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/xthet/go-morvo/services"
	"github.com/xthet/go-morvo/types"
	"github.com/xthet/go-morvo/utils"
)

type UserController struct {
	user_service *services.UserService
}

// if the controller is to be called outside here
func User(user_service *services.UserService) UserController {
	return UserController{
		user_service: user_service,
	}
}

func (c UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	token, err := c.user_service.LoginUser(&user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (c UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// payload validation
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.user_service.RegisterUser(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}
