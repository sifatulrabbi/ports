package controllers

import (
	"net"
	"net/http"
	"strings"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
	"github.com/sifatulrabbi/ports/pkg/validators"
	"golang.org/x/crypto/bcrypt"
)

type signInPayload struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

// Handle register request.
func Register(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.Response{}
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
	res := utils.Response{}
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		res.Data = nil
		res.Message = "Passwords don't match"
		res.BadRequest(w)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr) // Get the ip address
	session, err := services.CreateSession(user, ip)
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

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	token := r.Header.Get("Authorization")
	if token == "" {
		res.Message = "No refresh token found"
		res.BadRequest(w)
		return
	}
	refreshToken := strings.Split(token, " ")[1]
	accessToken, err := services.CreateAccessToken(refreshToken)
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Message = "Access token generated"
	res.Data = accessToken
	res.Ok(w)
}

func SignOut(w http.ResponseWriter, r *http.Request) error {
	return nil
}
