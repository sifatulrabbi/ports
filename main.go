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
	// IO server
	server := io.NewServer(ioOptions)

	server.OnConnect("/", func(s io.Conn) error {
		s.SetContext("")
		log.Println("connected: ", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s io.Conn, msg string) error {
		log.Println("notice: ", msg)
		s.Emit("reply", "have "+msg)
		return nil
	})

	server.OnError("/", func(s io.Conn, err error) {
		log.Fatalln("met error: ", s.ID(), err)
	})

	server.OnDisconnect("/", func(s io.Conn, reason string) {
		log.Println(s.ID(), "disconnected")
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalln("Socket.IO error: ", err)
		}
	}()
	defer server.Close()

	r.HandleFunc("/hello", controllers.Hello)
	r.Handle("/socket.io/", server)
	log.Printf("Starting the server on port %v\n", configs.Globals.PORT)
	if err := http.ListenAndServe(configs.Globals.PORT, r); err != nil {
		log.Fatalln(err)
	}

}
