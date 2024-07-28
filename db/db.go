package db

import (
	"context"
	"fmt"
	"log"

	"github.com/xthet/go-morvo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	client_options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), client_options)

	if err != nil {log.Fatal(); return nil, err}
	err = test_DB(client)
	if err != nil {log.Fatal(); return nil, err}
	fmt.Println("CONNECTED TO DATABASE SUCCESSFULLY!!")
	models.Collections(client)

	return client, nil
}

func test_DB(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil { 
		log.Fatal(err)
		return err
	}
	return nil
}