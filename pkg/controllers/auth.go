package controllers

import (
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
	"github.com/sifatulrabbi/ports/pkg/validators"
)

type signInPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	res := utils.CustomResponse{}
	p := signInPayload{}
	err := utils.BodyParser(r, &p)
	if err != nil {
		res.Data = nil
		res.Message = "Unable to parse request body"
		res.BadRequest(w)
		return
	}
	user, err := services.FindUserByUsername(p.Username)
	if err != nil {
		res.Data = nil
		res.Message = err.Error()
		res.NotFound(w)
		return
	}
	session, err := services.CreateSession(r, user)
	if err != nil {
		res.Data = user
		res.Message = "Unable to create session for the user"
		res.Internal(w)
		return
	}
	res.Message = "Login successful"
	res.Data = map[string]string{"refreshToken": session.RefreshToken, "username": user.Username, "email": user.Email}
	res.Ok(w)
}
