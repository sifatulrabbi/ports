package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
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
	if p.Username == "" || p.Password == "" {
		res.Message = "Invalid request body"
		res.BadRequest(w)
		return
	}

	u, err := services.FindUserByUsername(p.Username)
	if err != nil {
		res.Message = err.Error()
		return
	}

	// compare the passwords
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p.Password))
	if err != nil {
		res.Message = "Invalid credentials"
		res.Forbidden(w)
		return
	}

	// create a refresh token
	hasher := sha256.New()
	hasher.Write([]byte(uuid.NewString()))
	token := hasher.Sum(nil)
	// create a cookie handler and set the refresh token to the browser.
	cookieHandler := securecookie.New([]byte(token), securecookie.GenerateRandomKey(32))
	encoded, err := cookieHandler.Encode("PSID", token)
	if err != nil {
		res.Message = "Unable to set cookies"
		res.BadRequest(w)
		return
	}
	cookie := http.Cookie{
		Name:     "PSID",
		Value:    encoded,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
	}
	http.SetCookie(w, &cookie)

	// create jwt token for the client.
	claims := models.AuthTokenClaims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			Issuer:    "ports-app",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := jwtToken.SignedString([]byte(token))
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Data = AuthResponse{AccessToken: accessToken}
	res.Message = "user created"
	res.Ok(w)
}

func logout(w http.ResponseWriter, r *http.Request) {

}
