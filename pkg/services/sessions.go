package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	// jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/golang-jwt/jwt/v4"
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
	_, err := hasher.Write([]byte(u.Username + u.Email + u.Password + fmt.Sprint(time.Now().Unix())))
	str := hex.EncodeToString(hasher.Sum(nil))
	return str, err
}

// Create a login session for an user.
func CreateSession(u models.User, ip string) (models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s := models.Session{}
	// Get a refresh token.
	rToken, err := createHash(u)
	if err != nil {
		return s, err
	}
	// Update the session document.
	s.ID = primitive.NewObjectID()
	s.RefreshToken = rToken
	s.UserID = u.ID
	s.Username = u.Username
	s.Email = u.Email
	s.CreatedAt = time.Duration(time.Now().Unix())
	s.IP = ip
	// Save on the database.
	_, err = sessionsCollection().InsertOne(ctx, s)
	return s, err
}

func CreateAccessToken(refreshToken string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := sessionsCollection().FindOne(ctx, &bson.D{{Key: "refresh_token", Value: refreshToken}})
	if res.Err() != nil {
		return "", res.Err()
	}
	s := models.Session{}
	res.Decode(&s)
	exp := jwt.NewNumericDate(time.Now().Add(5 * time.Minute))
	authClaims := models.AuthTokenClaims{
		UserID:   s.UserID,
		Username: s.Username,
		Email:    s.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
			Issuer:    "ports-test",
		},
	}
	parseAuth := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	authToken, err := parseAuth.SignedString([]byte(s.RefreshToken))
	if err != nil {
		return "", err
	}
	accessClaims := models.AccessTokenClaims{
		ID:        s.ID,
		AuthToken: authToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
			Issuer:    "ports-test",
		},
	}
	parseAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := parseAccess.SignedString([]byte(configs.Globals.JWT_SECRET))
	return accessToken, err
}

func RemoveSession(id primitive.ObjectID, ip string, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "username", Value: username}},
				bson.D{{Key: "ip", Value: ip}},
			},
		},
		{Key: "$or",
			Value: bson.A{
				bson.D{{Key: "id", Value: id}},
			},
		},
	}
	_, err := sessionsCollection().DeleteMany(ctx, filter)
	return err
}
