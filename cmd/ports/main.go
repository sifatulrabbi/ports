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

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(w, r)
	})

	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, c.Handler(r)); err != nil {
		log.Fatalln(err)
	}
}
