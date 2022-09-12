package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
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

func HelloGET(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	res.Message = "GET request accepted"
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	rIp := r.Header.Get("X-FORWARDED-FOR")
	res.Data = map[string]string{"ip": ip, "realIp": rIp}
	res.Ok(w)
}

func HelloPOST(w http.ResponseWriter, r *http.Request) {
	var payload interface{}
	res := utils.Response{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Println(err)
		res.Message = "Unable to parse the body"
		res.Data = nil
		res.BadRequest(w)
		return
	}
	res.Message = "POST request accepted"
	res.Data = payload
	res.Ok(w)
}

func TestMongoDB(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	// Extract request body.
	var payload interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		res.Data = nil
		res.BadRequest(w)
		return
	}

	// Save data to database.
	collection := configs.GetCollection(configs.MongoClient, "users")
	insRes, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Md Sifatul Islam Rabbi"},
		{Key: "age", Value: 20},
		{Key: "occupation", Value: "Full stack developer"},
	})
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		res.Data = nil
		res.BadRequest(w)
		return
	}

	// Send success response
	res.Message = "Data saved on mongodb"
	res.Data = map[string]interface{}{"id": insRes.InsertedID}
	res.Created(w)
}
