package models

import (
	"context"
	"time"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Username  string             `json:"username,omitempty" validate:"required"`
	Password  string             `json:"password,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required"`
	Fullname  string             `json:"fullname,omitempty" validate:"required"`
	AvatarURL string             `json:"avatarURL,omitempty" validate:"required"`
	Bio       string             `json:"bio,omitempty" validate:"required"`
}

// Save the user.
func (u *User) Save() (*mongo.InsertOneResult, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	u.ID = primitive.NewObjectID()
	res, err := configs.GetCollection(configs.MongoClient, "users").InsertOne(ctx, u)
	return res, err
}

func (u *User) Update() error {
	return nil
}
