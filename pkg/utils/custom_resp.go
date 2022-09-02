package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResp struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Default internal error response body.
var internalErr = CustomResp{
	Success:    false,
	StatusCode: 500,
	Message:    "Internal server error",
	Data:       nil,
}

func (res CustomResp) Send(w http.ResponseWriter) error {
	var err error
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(res)
	if err == nil {
		w.Write(b)
	} else {
		log.Print(err)
		b, _ = json.Marshal(internalErr)
		w.Write(b)
	}
	return err
}

func (res *CustomResp) Ok(data interface{}, msg string) {
	res.Success = true
	res.StatusCode = http.StatusOK
	if msg == "" {
		res.Message = "Ok"
	}
	res.Data = data
}

func (res *CustomResp) BadRequest(data interface{}, msg string) {
	res.Success = false
	res.StatusCode = http.StatusBadRequest
	if msg == "" {
		res.Message = "Bad request"
	} else {
		res.Message = msg
	}
}

func (res *CustomResp) Created(data interface{}, msg string) {
	res.Success = true
	res.StatusCode = http.StatusCreated
	if msg == "" {
		res.Message = "Created"
	} else {
		res.Message = msg
	}
}

func (res *CustomResp) NotFound(data interface{}, msg string) {
	res.Success = false
	res.StatusCode = http.StatusNotFound
	if msg == "" {
		res.Message = "Not found"
	} else {
		res.Message = msg
	}
}

func (res *CustomResp) Unauthorized(data interface{}, msg string) {
	res.Success = false
	res.StatusCode = http.StatusUnauthorized
	if msg == "" {
		res.Message = "Unauthorized"
	} else {
		res.Message = msg
	}
}

func (res *CustomResp) Forbidden(data interface{}, msg string) {
	res.Success = false
	res.StatusCode = http.StatusForbidden
	if msg == "" {
		res.Message = "Forbidden"
	} else {
		res.Message = msg
	}
}
