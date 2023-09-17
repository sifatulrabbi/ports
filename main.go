package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sifatulrabbi/ports/api"
	"github.com/sifatulrabbi/ports/services"
)

var GO_ENV = os.Getenv("GO_ENV") // by default is going to be empty ""

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

	if GO_ENV != "production" {
		godotenv.Load(".env")
	}
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

	Router = setupRouter()
	api.RegisterUsersHandlers(Router, usersService)

	if err = Router.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://finetrack.sifatul.com", "http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Access-Control-Request-Headers",
		"Access-Control-Request-Method",
		"Authorization",
		"Connection",
		"Content-Type",
		"Cookie",
		"Date",
		"If-Modified-Since",
		"If-None-Match",
		"Origin",
		"Referrer",
		"User-Agent",
		"X-Requested-With",
	}
	r.Use(cors.New(corsConfig))

	return r
}
