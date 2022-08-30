package controllers

import (
	"log"

	io "github.com/googollee/go-socket.io"
)

type IOUserConn struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func RegisterHandlers(s *io.Server) {
	s.OnConnect("/", func(s io.Conn) error {
		s.SetContext("")
		log.Println("connected: ", s.ID())
		return nil
	})

	s.OnEvent("/", "notice", func(s io.Conn, msg string) error {
		log.Println("notice: ", msg)
		s.Emit("reply", "have "+msg)
		return nil
	})

	s.OnError("/", func(s io.Conn, err error) {
		log.Fatalln("met error: ", s.ID(), err)
	})

	s.OnDisconnect("/", func(s io.Conn, reason string) {
		log.Println(s.ID(), "disconnected")
	})

	s.OnEvent("/users", "authorize", UserConn)
}

func UserConn(s io.Conn, cred IOUserConn) error {
	var err error
	return err
}
