package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/controllers"
	"github.com/sifatulrabbi/ports/pkg/socket"
)

func main() {
	router := mux.NewRouter()
	configs.LoadENVs()
	configs.ConnectDB()
	// Register all the routes
	controllers.RegisterRoutes(router)
	// Handling CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(w, r)
	})

	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, c.Handler(router)); err != nil {
		log.Fatalln(err)
	}
}
