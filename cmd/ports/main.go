package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/controllers"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func main() {
	var (
		r = mux.NewRouter()
	)
	configs.LoadENVs()
	configs.ConnectDB()

	r.HandleFunc("/hello", controllers.HelloGET).Methods("GET")
	r.HandleFunc("/hello", controllers.TestMongoDB).Methods("POST")
	r.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/signin", controllers.SignIn).Methods("POST")
	r.HandleFunc("/api/v1/auth/accesstoken", controllers.GetAccessToken).Methods("GET")
	r.HandleFunc("/api/v1/users/{username}", utils.AuthGuard(controllers.GetUserByUsername)).Methods("GET")

	originOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"x-custom-header"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, handlers.CORS(originOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatalln(err)
	}
}
