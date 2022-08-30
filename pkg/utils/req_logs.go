package utils

import (
	"log"
	"net/http"
)

func LogReq(r *http.Request) {
	log.Printf("new req arrived, path: %v, method: %v\n", r.URL.Path, r.Method)
}
