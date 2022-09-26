package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

// get the user with the user name and return the profile expect password
func getProfile(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	queries := mux.Vars(r)
	u, err := services.FindUserByUsername(queries["username"])
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	u.Password = ""
	res.Data = u
	res.Message = "user found"
	res.Ok(w)
}
