package services

import (
	"HomeCloud/src/database"
	"HomeCloud/src/database/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func FindByUsername(data models.User) (error, models.User) {
	var user models.User

	err := database.DB.Collection("users").FindOne(context.Background(), bson.M{"username": data.Username}).Decode(&user)

	if err != nil {
		return err, user
	}

	return nil, user
}

func CreateUser(data models.User) error {
	_, err := database.DB.Collection("users").InsertOne(context.Background(), data)

	if err != nil {
		return err
	}

	return nil
}
