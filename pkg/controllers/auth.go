package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/utils"
	"github.com/sifatulrabbi/ports/pkg/validators"
)

// Handle register request.
func Register(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	valid := validators.RegisterPayload(&user)
	if !valid {
		b, _ := json.Marshal(map[string]interface{}{"success": false, "message": "Invalid payload"})
		w.Write(b)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

// Handle sign in.
func SignIn(w http.ResponseWriter, r *http.Request) {
	utils.LogReq(r)
}
