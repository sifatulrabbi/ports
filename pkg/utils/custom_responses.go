package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Default internal error response body.
var internalErr = CustomResponse{
	Success:    false,
	StatusCode: 500,
	Message:    "Internal server error",
	Data:       nil,
}

func (res CustomResponse) Send(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
		b, _ = json.Marshal(internalErr)
		w.Write(b)
		return
	}
	w.WriteHeader(res.StatusCode)
	w.Write(b)
}

func (res *CustomResponse) Ok(w http.ResponseWriter) {
	res.Success = true
	res.StatusCode = http.StatusOK
	if res.Message == "" {
		res.Message = "Ok"
	}
	res.Send(w)
}

func (res *CustomResponse) BadRequest(w http.ResponseWriter) {
	res.Success = false
	res.StatusCode = http.StatusBadRequest
	if res.Message == "" {
		res.Message = "Bad request"
	}
	res.Send(w)
}

func (res *CustomResponse) Created(w http.ResponseWriter) {
	res.Success = true
	res.StatusCode = http.StatusCreated
	if res.Message == "" {
		res.Message = "Created"
	}
	res.Send(w)
}

func (res *CustomResponse) NotFound(w http.ResponseWriter) {
	res.Success = false
	res.StatusCode = http.StatusNotFound
	if res.Message == "" {
		res.Message = "Not found"
	}
	res.Send(w)
}

func (res *CustomResponse) Unauthorized(w http.ResponseWriter) {
	res.Success = false
	res.StatusCode = http.StatusUnauthorized
	if res.Message == "" {
		res.Message = "Unauthorized"
	}
	res.Send(w)
}

func (res *CustomResponse) Forbidden(w http.ResponseWriter) {
	res.Success = false
	res.StatusCode = http.StatusForbidden
	if res.Message == "" {
		res.Message = "Forbidden"
	}
	res.Send(w)
}

func (res *CustomResponse) Internal(w http.ResponseWriter) {
	res.Success = false
	res.StatusCode = http.StatusInternalServerError
	if res.Message == "" {
		res.Message = "Internal server error"
	}
	res.Send(w)
}
