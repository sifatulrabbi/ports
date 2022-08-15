package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Fatalln("error while decoding json body", err)
	}
	log.Println(u)

	jsonResp, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(jsonResp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}
