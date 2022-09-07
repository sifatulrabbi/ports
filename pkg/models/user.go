package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"id" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	Email     string             `bson:"email" json:"email"`
	Fullname  string             `bson:"fullname" json:"fullname"`
	AvatarURL string             `bson:"avatar_url" json:"avatar_url"`
	Bio       string             `bson:"bio" json:"bio"`
}
