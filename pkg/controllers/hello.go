package controllers

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello world"))
}
