package controllers

import (
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
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
	user, err := services.CreateUser(user)
	if err != nil {
		res.Message = err.Error()
		res.Data = nil
		res.BadRequest(w)
		return
	}
	if err != nil {
		res.Message = "Unable to generate session"
		res.Data = nil
		res.BadRequest(w)
		return
	}
	res.Message = "User created"
	res.Data = user
	res.Created(w)
}

// Handle sign in.
func SignIn(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
}
