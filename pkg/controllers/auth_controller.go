package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

type registerBody struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
	Fullname        string `json:"fullname"`
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

// register controller, verifies and creates a new user account.
func register(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	p := registerBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		res.Message = "invalid request body"
		res.BadRequest(w)
		return
	}
	err = r.Body.Close()
	if err != nil {
		res.Message = "invalid request body"
		res.BadRequest(w)
		return
	}
	if p.Username == "" || p.Password == "" || p.ConfirmPassword == "" {
		res.Message = "invalid request body"
		res.BadRequest(w)
		return
	}
	if p.Password != p.ConfirmPassword {
		res.Message = "passwords don't match"
		res.BadRequest(w)
		return
	}
	// verify the authenticity of the username
	if prevUser, err := services.FindUserByUsername(p.Username); err == nil {
		if prevUser.Username == p.Username {
			res.Message = "username already in use"
			res.BadRequest(w)
			return
		}
	}

	// create the user account.
	_, err = services.CreateUser(
		models.User{
			Username: p.Username,
			Password: p.Password,
			Email:    p.Email,
			Fullname: p.Fullname,
		})
	if err != nil {
		fmt.Println(err)
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}

	res.Data = true
	res.Message = "user created"
	res.Ok(w)
}

// login controller, to verify user's password ans create new session for the user.
func login(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	p := loginBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
		res.Message = "invalid request body"
		res.BadRequest(w)
		return
	}
	err = r.Body.Close()
	if err != nil {
		fmt.Println(err)
		res.Message = "invalid request body"
		res.BadRequest(w)
		return
	}
	if p.Username == "" || p.Password == "" {
		fmt.Println(err)
		res.Message = "Invalid request body"
		res.BadRequest(w)
		return
	}

	u, err := services.FindUserByUsername(p.Username)
	if err != nil {
		fmt.Println(err)
		res.Message = err.Error()
		res.NotFound(w)
		return
	}

	// compare the passwords
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p.Password))
	if err != nil {
		fmt.Println(err)
		res.Message = "Invalid credentials"
		res.Forbidden(w)
		return
	}

	refreshToken, accessToken, err := services.CreateSession(r, &u)
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	cookie := http.Cookie{
		Name:  "PSID",
		Value: refreshToken,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	res.Data = AuthResponse{AccessToken: accessToken}
	res.Message = "user created"
	res.Ok(w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	// remove the refresh token cookie
	cookie := http.Cookie{
		Name:   "PSID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	res.Message = "successfully logged out"
	res.Ok(w)
}
