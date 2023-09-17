package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (r *HttpResponse) BadRequest(c *gin.Context) {
	r.StatusCode = http.StatusBadRequest
	if r.Message == "" {
		r.Message = "Bad request"
	}
	r.Success = false
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) Unauthorized(c *gin.Context) {
	r.StatusCode = http.StatusUnauthorized
	if r.Message == "" {
		r.Message = "Unauthorized"
	}
	r.Success = false
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) Forbidden(c *gin.Context) {
	r.StatusCode = http.StatusForbidden
	if r.Message == "" {
		r.Message = "Forbidden"
	}
	r.Success = false
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) Ok(c *gin.Context) {
	r.StatusCode = http.StatusOK
	if r.Message == "" {
		r.Message = "Request was successful"
	}
	r.Success = true
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) Created(c *gin.Context) {
	r.StatusCode = http.StatusCreated
	if r.Message == "" {
		r.Message = "Resource created successfully"
	}
	r.Success = true
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) NotFound(c *gin.Context) {
	r.StatusCode = http.StatusNotFound
	if r.Message == "" {
		r.Message = "Resource not found"
	}
	r.Success = false
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) InternalServerError(c *gin.Context) {
	r.StatusCode = http.StatusInternalServerError
	if r.Message == "" {
		r.Message = "Internal server error"
	}
	r.Success = false
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func (r *HttpResponse) New(c *gin.Context, statusCode int) {
	r.StatusCode = statusCode
	c.JSON(r.StatusCode, r)
	c.Abort()
}

func routeWrapper[Service any](s Service, handler func(c *gin.Context, s Service)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, s)
	}
}
