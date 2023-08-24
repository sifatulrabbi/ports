package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sifatulrabbi/ports/api"
	"github.com/sifatulrabbi/ports/services"
)

func main() {
	var (
		err      error
		port     string
		pgHost   string
		pgDbName string
		pgSSL    string
		db       *gorm.DB
		Router   *gin.Engine
	)

	godotenv.Load(".env")
	port = os.Getenv("PORT")
	pgHost = os.Getenv("PGHOST")
	pgDbName = os.Getenv("PGDBNAME")
	pgSSL = os.Getenv("PGSSL")
	if port == "" {
		log.Fatalln("PORT env not found")
	}

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=%s", pgHost, pgDbName, pgSSL)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to the db\nerror: %v\n", err)
	}

	usersService, err := services.NewUsersService(db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	Router = gin.Default()
	api.RegisterUsersHandlers(Router, usersService)

	if err = Router.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
