package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/auth/register", controllerWrapper(register)).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", controllerWrapper(login)).Methods("POST")
	r.HandleFunc("/api/v1/auth/logout", controllerWrapper(logout)).Methods("POST")
}

func controllerWrapper(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v:%v\n", r.Method, r.URL)
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		next(w, r)
	}
}
