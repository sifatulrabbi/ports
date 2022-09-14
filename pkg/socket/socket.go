package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	ID      int
	Payload []byte
}

const (
	READ_SIZE  = 1024
	WRITE_SIZE = 1024
)

var (
	readPump  = make(chan Message, 256)
	writePump = make(chan Message, 256)
	clients   = map[int]*websocket.Conn{}
	lastId    = 0
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  READ_SIZE,
	WriteBufferSize: WRITE_SIZE,
}

func reader() {
	for msg := range readPump {
		log.Printf("Message from: %v", msg.ID)
	}
}

func writer() {
	for msg := range writePump {

	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("Client successfully connected")
	go reader()

	currId := lastId
	lastId = lastId + 1
	clients[currId] = conn
	for {
		_, p, err := conn.ReadMessage()
		msg := Message{ID: currId, Payload: p}
		if err != nil {
			log.Println(err)
			return
		}
		readPump <- msg
	}
}
