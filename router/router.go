package router

import (
	"net/http"

	"github.com/xthet/go-morvo/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(client *mongo.Client) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /", greet)

	TodoRoutes(router, client)
	UserRoutes(router, client)

	sub_router := http.NewServeMux()
	sub_router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))	

	return sub_router
}

func greet(w http.ResponseWriter, r *http.Request){
	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg":"Hello World"})
}
