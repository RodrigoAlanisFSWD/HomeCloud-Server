package services

import (
	"HomeCloud/src/database"
	"HomeCloud/src/database/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindByUsername(data models.User) (error, models.User) {
	var user models.User

	err := database.DB.Collection("users").FindOne(context.Background(), bson.M{"username": data.Username}).Decode(&user)

	if err != nil {
		return err, user
	}

	return nil, user
}

func FindById(id primitive.ObjectID) (models.User, error) {
	var user models.User

	err := database.DB.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return user, nil
	}

	return user, nil
}

func CreateUser(data models.User) error {
	data.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("users").InsertOne(context.Background(), data)

	if err != nil {
		return err
	}

	return nil
}

func UpdateAvatar(id string, name string) error {
	var user models.User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"avatar": "http://localhost:5000/assets/" + name}}

	err = database.DB.Collection("users").FindOneAndUpdate(context.Background(), bson.M{"_id": objectId}, update).Decode(&user)

	if err != nil {
		return err
	}

	return nil
}
