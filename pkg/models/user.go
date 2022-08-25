package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatarURL"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
