package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `json:"id"`
	CreatedAt    time.Duration      `json:"createdAt"`
	UserID       primitive.ObjectID `json:"userId"`
	Username     string             `json:"username"`
	RefreshToken string             `json:"refreshToken"`
	IP           string             `json:"ip"`
}
