package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/controllers"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/hello", controllers.HelloGET).Methods("GET")
	r.HandleFunc("/hello", controllers.TestMongoDB).Methods("POST")
	// Auth routers
	r.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/signin", controllers.SignIn).Methods("POST")
	r.HandleFunc("/api/v1/auth/accesstoken", controllers.GetAccessToken).Methods("GET")
	r.HandleFunc("/api/v1/users/{username}", utils.AuthGuard(controllers.GetUserByUsername)).Methods("GET")
	// Directory management routers
	r.HandleFunc("/api/v1/directories", controllers.GetDirNames).Methods("GET")
	r.HandleFunc(`/api/v1/directories/{path:[a-zA-Z0-9/_\-\.\?]+}`, controllers.GetSubDirs).Methods("GET")
	r.HandleFunc(`/api/v1/files/{path:[a-zA-Z0-9/_\-\.\?]+}`, controllers.GetAFile).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	configs.LoadENVs()
	configs.ConnectDB()
	// Register all the routes
	registerRoutes(r)
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
