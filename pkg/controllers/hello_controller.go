package controllers

import (
	"net"
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/utils"
)

type BasicMsg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type BasicResp struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func helloGET(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	res.Message = "GET request accepted"
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	rIp := r.Header.Get("X-FORWARDED-FOR")
	res.Data = map[string]string{"ip": ip, "realIp": rIp}
	res.Ok(w)
}
