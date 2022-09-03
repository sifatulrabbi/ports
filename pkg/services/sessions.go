package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/models"
)

// Get sessions collection.
func sessionsCollection() *mongo.Collection {
	return configs.GetCollection(configs.MongoClient, "sessions")
}

// Create a hash with use user's credentials and time stamp.
func createHash(u models.User) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(u.Username + u.Email + u.Password))
	str := hex.EncodeToString(hasher.Sum(nil))
	return str, err
}

// Create a login session for an user.
func CreateSession(r *http.Request, u models.User) (models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s := models.Session{}
	// Get a refresh token.
	rToken, err := createHash(u)
	if err != nil {
		return s, err
	}
	// Update the session document.
	s.RefreshToken = rToken
	s.ID = primitive.NewObjectID()
	s.UserID = u.ID
	s.Username = u.Username
	s.CreatedAt = time.Duration(time.Now().Unix())
	// Save on the database.
	_, err = sessionsCollection().InsertOne(ctx, s)
	return s, err
}
