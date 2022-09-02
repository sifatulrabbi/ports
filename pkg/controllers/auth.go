package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/utils"
	"github.com/sifatulrabbi/ports/pkg/validators"
)

// Handle register request.
func Register(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.CustomResponse{}
	user := models.User{}
	utils.BodyParser(r, &user)

	// Validate the payload.
	valid := validators.RegisterPayload(&user)
	if !valid {
		res.Message = "Invalid request payload"
		res.Data = nil
		res.BadRequest(w)
		return
	}

	// Save the user.
	_, err := user.Save()
	fmt.Println(user)
	if err != nil {
		res.Message = err.Error()
		res.Data = nil
		res.BadRequest(w)
		return
	}
	res.Message = "User created"
	res.Data = user
	log.Println(res)
	res.Created(w)
}

// Handle sign in.
func SignIn(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
}
