package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResp struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Default internal error response body.
var internalErr = CustomResp{
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

func (res *CustomResp) Ok(w http.ResponseWriter) error {
	res.StatusCode = http.StatusOK
	if res.Message == "" {
		res.Message = "Ok"
	}
	err := res.Send(w)
	return err
}

func (res *CustomResp) BadRequest(w http.ResponseWriter) error {
	res.StatusCode = http.StatusBadRequest
	if res.Message == "" {
		res.Message = "Bad request"
	}
	err := res.Send(w)
	return err
}

func (res *CustomResp) Created(w http.ResponseWriter) error {
	res.StatusCode = http.StatusCreated
	if res.Message == "" {
		res.Message = "Created"
	}
	err := res.Send(w)
	return err
}

func (res *CustomResp) NotFound(w http.ResponseWriter) error {
	res.StatusCode = http.StatusNotFound
	if res.Message == "" {
		res.Message = "Not found"
	}
	err := res.Send(w)
	return err
}

func (res *CustomResp) Unauthorized(w http.ResponseWriter) error {
	res.StatusCode = http.StatusUnauthorized
	if res.Message == "" {
		res.Message = "Unauthorized"
	}
	err := res.Send(w)
	return err
}

func (res *CustomResp) Forbidden(w http.ResponseWriter) error {
	res.StatusCode = http.StatusForbidden
	if res.Message == "" {
		res.Message = "Forbidden"
	}
	err := res.Send(w)
	return err
}
