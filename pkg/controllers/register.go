package controllers

import (
	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func Register(r *mux.Router) {
	r.HandleFunc("/hello", helloGET).Methods("GET")
	// Auth routers
	r.HandleFunc("/api/v1/auth/register", register).Methods("POST")
	r.HandleFunc("/api/v1/auth/signin", signIn).Methods("POST")
	r.HandleFunc("/api/v1/auth/accesstoken", getAccessToken).Methods("GET")
	r.HandleFunc("/api/v1/users/{username}", utils.AuthGuard(GetUserByUsername)).Methods("GET")
	// Directory management routers
	r.HandleFunc("/api/v1/directories", getDirNames).Methods("GET")
	r.HandleFunc(`/api/v1/directories/{path:[a-zA-Z0-9/_\-\.\?]+}`, getSubDirs).Methods("GET")
	r.HandleFunc(`/api/v1/files/{path:[a-zA-Z0-9/_\-\.\?]+}`, getAFile).Methods("GET")
}
