package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessTokenClaims struct {
	ID        primitive.ObjectID `json:"id"`         // id of the associated refresh token
	AuthToken string             `json:"auth_token"` // the actual token that contains user information
	jwt.RegisteredClaims
}

type AuthTokenClaims struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Username string             `json:"username"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	RefreshToken string             `bosn:"refresh_token"` // Session's refresh token
	IP           string             `bson:"ip"`            // The document will also store the IP address to ensure security
	Iat          int64              `bson:"iat"`           // Session initiation date
	UserID       primitive.ObjectID `json:"user_id"`
	Username     string             `bson:"username"`
	jwt.RegisteredClaims
}
