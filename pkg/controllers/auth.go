package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/sifatulrabbi/ports/pkg/utils"
)

type TestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	decoder := json.NewDecoder(r.Body)
	tb := TestBody{}
	err := decoder.Decode(&tb)
	if err != nil {
		log.Fatalln(err)
	}
	// hash password.
	hp, err := bcrypt.GenerateFromPassword([]byte(tb.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	tb.Password = string(hp)

	// send a response.
	res := utils.CustomResp{
		Data:    tb,
		Message: "User created",
	}
	res.Ok(w)
}
