package main

import (
	"log"
	"net/http"

	io "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/gorilla/mux"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/controllers"
)

func allowOrigin(r *http.Request) bool {
	return true
}

func main() {
	var (
		r = mux.NewRouter()
	)
	configs.LoadENVs()
	// Socket io configs to prevent CORS error.
	ioOptions := &engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{CheckOrigin: allowOrigin},
			&websocket.Transport{CheckOrigin: allowOrigin},
		},
	}
	// IO server.
	server := io.NewServer(ioOptions)
	r.Handle("/socket.io/", server)
	// Register routes.
	r.HandleFunc("/hello", controllers.HelloGET).Methods("GET")
	r.HandleFunc("/hello", controllers.TestMongoDB).Methods("POST")
	r.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/signin", controllers.SignIn).Methods("POST")

	configs.ConnectDB()
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalln("Socket.IO error: ", err)
		}
	}()
	defer server.Close()

	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, r); err != nil {
		log.Fatalln(err)
	}
}
