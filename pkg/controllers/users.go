package controllers

import (
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
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

func GetUserByUsername(w http.ResponseWriter, r *http.Request, session *models.Session) {
	res := utils.Response{}
	user, err := services.FindUserById(session.UserID)
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Message = "User found"
	res.Data = user
	res.Ok(w)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// res := utils.CustomResponse{}
	// user, err := services.FindUserById(session.UserID)
	// if err != nil {
	// 	res.Message = err.Error()
	// 	res.BadRequest(w)
	// 	return
	// }
	// res.Message = "User found"
	// res.Data = user
	// res.Ok(w)
}
