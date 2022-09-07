package models

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `bson:"id" json:"id"`
	IP           string             `bson:"ip" json:"ip"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	Username     string             `bson:"username" json:"username"`
	Email        string             `bson:"email" json:"email"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt    time.Duration      `bson:"created_at" json:"created_at"`
}

type AccessTokenClaims struct {
	ID        primitive.ObjectID `json:"id"`         // id of the associated refresh token
	AuthToken string             `json:"auth_token"` // the actual token that contains user information
	jwt.RegisteredClaims
}

type AuthTokenClaims struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}
