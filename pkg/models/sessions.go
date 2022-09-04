package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID        string             `json:"id"`
	CreatedAt time.Duration      `json:"createdAt"`
	UserID    primitive.ObjectID `json:"userId"`
	Username  string             `json:"username"`
	IP        string             `json:"ip"`
}
