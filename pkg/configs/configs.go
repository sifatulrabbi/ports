package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENVs struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_NAME     string
	DB_URI      string
	JWT_SECRET  string
}

var Globals ENVs

func LoadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No file found, filename: .env")
	}

	PORT := os.Getenv("PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_URI := os.Getenv("MONGODB_URL")
	DB_NAME := os.Getenv("DB_NAME")
	JWT_SECRET := os.Getenv("JWT_SECRET")

	// Check if all the env vars are present or not.
	if PORT == "" || DB_USERNAME == "" || DB_PASSWORD == "" || DB_URI == "" || DB_NAME == "" {
		log.Fatalln("not enough env vars")
	}

	Globals.PORT = ":" + PORT
	Globals.DB_USERNAME = DB_USERNAME
	Globals.DB_PASSWORD = DB_PASSWORD
	Globals.DB_URI = DB_URI
	Globals.DB_NAME = DB_NAME
	Globals.JWT_SECRET = JWT_SECRET
	if PORT == "" {
		log.Fatal("env var PORT not found")
	}
}
