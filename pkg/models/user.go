package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
