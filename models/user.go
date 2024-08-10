package models

import (
	"context"

	"github.com/xthet/go-morvo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	collection *mongo.Collection
}

func User(collection *mongo.Collection) *UserCollection {
	return &UserCollection{
		collection: collection,
	}
}

func (c UserCollection) GetUserByID(userID primitive.ObjectID) (*types.User, error) {
	filter := bson.M{"id": userID}
	var user = new(types.User)
	err := c.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserCollection) GetUserByEmail(email string) (*types.User, error) {
	filter := bson.M{"email": email}
	var user = new(types.User)
	err := c.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserCollection) CreateUser(user types.User) error {
	_, err := c.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
