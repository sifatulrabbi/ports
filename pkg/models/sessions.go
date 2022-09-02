package models

import (
	"time"
)

type Sessions struct {
	CreatedAt    time.Duration
	ExpiresAt    time.Duration
	UserID       string
	RefreshToken string
	IP           string
}
