package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// Create connection with mongodb.
func ConnectDB() *mongo.Client {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(Globals.DB_URI).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Mongodb connection failed")
		log.Fatal(err)
	}
	MongoClient = client
	return client
}

// Get a mongo db collection
func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Println("Mongodb connection failed")
		log.Fatal(err)
	}
	return client.Database(Globals.DB_NAME).Collection(colName)
}
