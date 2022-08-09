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
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	basicMsg := &BasicMsg{}
	err := utils.DecodeJson(w, r, basicMsg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(basicMsg)

	rb, _ := json.Marshal(&BasicResp{Status: http.StatusOK, Success: true, Message: basicMsg.Message})
	if _, err := w.Write(rb); err != nil {
		log.Fatal(err)
	}

}
