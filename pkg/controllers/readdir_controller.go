package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func GetDirNames(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.Response{}
	dirNames, err := services.GetHomeDir()
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Data = dirNames
	res.Message = "Directories found"
	res.Ok(w)
}

func GetDirInformation(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.Response{}
	params := mux.Vars(r)
	info, err := services.GetDirInfo(params["dirname"])
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Message = "Directory's information"
	res.Data = info
	res.Ok(w)
}

func GetSubDirs(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.Response{}
	params := mux.Vars(r)
	result, err := services.GetDirInfo(params["path"])
	if err != nil {
		res.Message = err.Error()
		res.Data = result
		res.BadRequest(w)
		return
	}
	res.Data = result
	res.Ok(w)
}

func GetAFile(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	res := utils.Response{}
	params := mux.Vars(r)
	content, err := services.GetFile(params["path"])
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(w)
		return
	}
	res.Data = content
	res.SendFile(w)
}
