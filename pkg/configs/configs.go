package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENVs struct {
	PORT string
}

var Globals ENVs

func LoadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No file found, filename: .env")
	}

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		log.Fatal("env var PORT not found")
	}
	Globals.PORT = ":" + PORT
}
