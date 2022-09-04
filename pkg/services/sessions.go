package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/models"
)

type jwtClaims struct {
	Username string             `json:"username"`
	Email    string             `json:"email"`
	UserID   primitive.ObjectID `json:"userId"`
	jwt.RegisteredClaims
}

// Get sessions collection.
func sessionsCollection() *mongo.Collection {
	return configs.GetCollection(configs.MongoClient, "sessions")
}

// Create a hash with use user's credentials and time stamp.
func createHash(u models.User) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(u.Username + u.Email + u.Password + fmt.Sprint(time.Now().Unix())))
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
	s.ID = rToken
	s.UserID = u.ID
	s.Username = u.Username
	s.CreatedAt = time.Duration(time.Now().Unix())
	s.IP, _, _ = net.SplitHostPort(r.RemoteAddr)
	// Save on the database.
	_, err = sessionsCollection().InsertOne(ctx, s)
	return s, err
}

func CreateAccessToken(rToken string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := sessionsCollection().FindOne(ctx, bson.D{{Key: "id", Value: rToken}})
	session := models.Session{}
	res.Decode(&session)
	user, err := FindUserById(session.UserID)
	if err != nil {
		err = errors.New("no session found for the access token")
		return "", err
	}
	claims := jwtClaims{
		Username: user.Username,
		Email:    user.Email,
		UserID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			Issuer:    "ports-test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(session.ID))
	if err != nil {
		err = errors.New("unable to sign the jwt token")
	}
	return ss, err
}
