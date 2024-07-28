package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var Todos *mongo.Collection

func Collections(client *mongo.Client) {
	Todos = client.Database("morvo").Collection("todos")
}	