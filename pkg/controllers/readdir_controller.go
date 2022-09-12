package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/ports/pkg/services"
	"github.com/sifatulrabbi/ports/pkg/utils"
)

func getDirNames(w http.ResponseWriter, r *http.Request) {
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

func getSubDirs(w http.ResponseWriter, r *http.Request) {
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

func getAFile(w http.ResponseWriter, r *http.Request) {
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
