package db

import (
	"fmt"
	"log"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable",
		configs.Globals.DB_USERNAME,
		configs.Globals.DB_PASSWORD,
		configs.Globals.DB_NAME,
		configs.Globals.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("unable to connect to the database")
		log.Fatalln(err)
	}

	DB = db
}
