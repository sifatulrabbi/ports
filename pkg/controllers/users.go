package controllers

import (
	"net/http"

	// "github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	// u := &models.User{}
	// respErr := utils.DecodeJson(w, r, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}
