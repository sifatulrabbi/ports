package configs

import (
	"log"

	"github.com/sifatulrabbi/ports/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var Mongo = MongoDB{}

// Create connection with mongodb.
func ConnectDB() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(Globals.DB_URI).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(utils.GetCtx(10), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Mongo.Client = client
	Mongo.DB = client.Database(Globals.DB_NAME)
	return client
}
