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
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_URI := os.Getenv("DB_URI")

	// Check if all the env vars are present or not.
	if PORT == "" || DB_USERNAME == "" || DB_PASSWORD == "" || DB_NAME == "" || DB_URI == "" {
		log.Fatalln("not enough env vars")
	}

	Globals.PORT = ":" + PORT
	Globals.DB_USERNAME = DB_USERNAME
	Globals.DB_PASSWORD = DB_PASSWORD
	Globals.DB_PORT = DB_PORT
	Globals.DB_NAME = DB_NAME
	Globals.DB_URI = DB_URI
	if PORT == "" {
		log.Fatal("env var PORT not found")
	}
}
