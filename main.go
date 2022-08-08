package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var (
		r     = mux.NewRouter()
		errCh = make(chan error)
	)

	fmt.Printf("Starting the server on port %v\n", 8000)
	go func() {
		if err := http.ListenAndServe(":8000", r); err != nil {
			errCh <- err
		}
	}()

	log.Fatal(<-errCh)
}
