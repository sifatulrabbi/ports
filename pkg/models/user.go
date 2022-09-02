package models

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

type User struct {
	ID        interface{} `bson:"id" json:"id"`
	Username  string      `bson:"username" json:"username"`
	Password  string      `bson:"password" json:"password"`
	Email     string      `bson:"email" json:"email"`
	Fullname  string      `bson:"fullname" json:"fullname"`
	AvatarURL string      `bson:"avatarURL" json:"avatarURL"`
	Bio       string      `bson:"bio" json:"bio"`
}

// Get the database collection for users.
func (u User) GetCollection() *mongo.Collection {
	if configs.Mongo.DB == nil {
		return nil
	}
	return configs.Mongo.DB.Collection("users")
}

// Create an user.
func CreateUser(u User) (*mongo.InsertOneResult, error) {
	return u.GetCollection().InsertOne(utils.GetCtx(10), u)
}
