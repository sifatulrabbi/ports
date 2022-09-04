package services

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

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
	// Hash the password.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hashedPass)
	_, err = usersCollection().InsertOne(ctx, u) // Save the user.
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
func FindUserById(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	res := usersCollection().FindOne(ctx, filter)
	user := models.User{}
	err := res.Decode(&user)
	if err != nil {
		return user, err
	}
	if user.Username == "" {
		return user, errors.New("user not found")
	}
	return user, nil
}

// Find user by username
func FindUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	res := usersCollection().FindOne(ctx, filter)
	user := models.User{}
	err := res.Decode(&user)
	return user, err
}
