package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
