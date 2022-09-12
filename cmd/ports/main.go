package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/controllers"
)

func main() {
	r := mux.NewRouter()
	configs.LoadENVs()
	configs.ConnectDB()
	// Register all the routes
	controllers.Register(r)
	// Handling CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Authorization", "Host", "Accept", "User-Agent"},
	})

	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, c.Handler(r)); err != nil {
		log.Fatalln(err)
	}
}
