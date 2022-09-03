package services

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get the users collection
func usersCollection() *mongo.Collection {
	return configs.GetCollection(configs.MongoClient, "users")
}

// Create an user on the database.
func CreateUser(u models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u.ID = primitive.NewObjectID()
	_, err := usersCollection().InsertOne(ctx, u)
	if err != nil {
		u = models.User{}
	}
	return u, err
}

// Update an user document.
func UpdateUser(u models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := usersCollection().UpdateByID(ctx, u.ID, u)
	if err != nil {
		u = models.User{}
	}
	return u, err
}

// Remove an user from the database.
func RemoveUser(u models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := usersCollection().DeleteOne(ctx, u)
	if err != nil {
		return false, err
	}
	return true, err
}

// Find an user from the database.
func FindUser(u models.User) (models.User, error) {
	fUser := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := usersCollection().FindOne(ctx, u)
	res.Decode(&fUser)
	if fUser.Username == "" {
		fUser = models.User{}
		return fUser, errors.New("user not found")
	}
	return fUser, nil
}
