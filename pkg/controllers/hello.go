package controllers

import (
	"encoding/json"
	"log"
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

func Hello(w http.ResponseWriter, r *http.Request) {
	var resp BasicResp

	basicMsg := &BasicMsg{}
	err := utils.DecodeJson(w, r, basicMsg)
	if err != nil {
		resp = BasicResp{Success: false, Status: err.Status, Message: err.Error(), Data: nil}
	}
	resp = BasicResp{Status: 200, Success: true, Message: "Successful", Data: basicMsg}

	rb, _ := json.Marshal(resp)
	if _, err := w.Write(rb); err != nil {
		log.Fatalln(err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	var resp BasicResp
	ct := r.Header.Get("Content-Type")
	resp = BasicResp{
		Success: true,
		Status:  200,
		Message: "Your content type",
		Data:    map[string]string{"Content-Type": ct},
	}
	w.Header().Add("Content-Type", "application/json")
	rb, _ := json.Marshal(resp)
	if _, err := w.Write(rb); err != nil {
		log.Fatalln(err)
	}
}
