package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/xthet/go-morvo/db"
	"github.com/xthet/go-morvo/router"
)

func main() {
	if os.Getenv("ENV") != "production" {
		// Load the .env file if not in production
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	client, err := db.ConnectMongoDB(MONGODB_URI)
	if err != nil {
        log.Fatal("Cannot connect to database", err)
    }

	defer client.Disconnect(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("api listening on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router.Routes())
	
	if err != nil {
		log.Fatal(err)
	}
}