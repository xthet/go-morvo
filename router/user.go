package router

import (
	"net/http"

	ctrls "github.com/xthet/go-morvo/controllers"
	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/services"
	"go.mongodb.org/mongo-driver/mongo"
)


func UserRoutes(router *http.ServeMux, client *mongo.Client){
	user_controller := ctrls.User(services.User(models.User(client.	Database("morvo").Collection("users"))))

	// USER ROUTES
	router.HandleFunc("POST /login", user_controller.LoginUser)
	router.HandleFunc("POST /register", user_controller.RegisterUser)
}