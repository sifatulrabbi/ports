package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/sifatulrabbi/ports/pkg/utils"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/auth/register", controllerWrapper(register)).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", controllerWrapper(login)).Methods("POST")
	r.HandleFunc("/api/v1/auth/logout", controllerWrapper(logout)).Methods("POST", "GET", "DELETE")
	r.HandleFunc("/api/v1/users/{username}", controllerWrapper(AuthGuard(getProfile))).Methods("GET")
}

func controllerWrapper(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v:%v\n", r.Method, r.URL)
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		next(w, r)
	}
}

func AuthGuard(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := utils.Response{}
		// get the session cookie and verify it
		var cookie *http.Cookie
		for _, c := range r.Cookies() {
			if c.Name == "PSID" {
				cookie = c
				break
			}
		}
		if cookie == nil {
			res.Message = "please login first"
			res.Forbidden(w)
			return
		}
		// get the authorization token and verify it
		token := r.Header.Get("Authorization")
		if token == "" {
			res.Message = "authorization token not found"
			res.Forbidden(w)
			return
		}
		accessToken := strings.Split(token, " ")[1]
		if accessToken == "" {
			res.Message = "invalid authorization token"
			res.Forbidden(w)
			return
		}
		next(w, r)
	}
}
