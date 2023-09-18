package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sifatulrabbi/ports/services"
)

func getUserByEmail(c *gin.Context, s *services.UsersService) {
	res := HttpResponse{}
	email := c.Query("email")
	if email == "" {
		res.Message = "No email found in the request"
		res.BadRequest(c)
		return
	}

	user, err := s.GetByEmail(services.UserFilter{Email: email})
	if err != nil {
		res.Message = err.Error()
		if res.Message == "user not found" {
			res.NotFound(c)
		} else {
			res.BadRequest(c)
		}
		return
	}
	res.Data = user.JSON()
	res.Message = "User profile found"
	res.Ok(c)
}

func getUserById(c *gin.Context, s *services.UsersService) {
	res := HttpResponse{}
	user, err := s.GetOne(services.UserFilter{})
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(c)
		return
	}
	res.Message = "User found"
	res.Data = user
	res.Ok(c)
	c.Abort()
}

func createOneUser(c *gin.Context, s *services.UsersService) {
	res := HttpResponse{}
	payload := services.UserPayload{}
	if err := c.BindJSON(&payload); err != nil {
		res.Message = err.Error()
		res.BadRequest(c)
		return
	}
	user, err := s.CreateOne(payload)
	if err != nil {
		res.Message = err.Error()
		res.BadRequest(c)
		return
	}
	res.Message = "New user created"
	res.Data = user
	res.Created(c)
	c.Abort()
}

func updateOneUser(c *gin.Context, s *services.UsersService) {
	res := HttpResponse{}
	payload := services.UserPayload{}
	if err := c.BindJSON(&payload); err != nil {
		res.Message = err.Error()
		res.BadRequest(c)
		return
	}

	res.Message = "API not implemented yet."
	res.Data = payload
	res.Ok(c)
}

func deleteOneUser(c *gin.Context, s *services.UsersService) {

}

func getManyUsers(c *gin.Context, s *services.UsersService) {

}

func RegisterUsersHandlers(r *gin.Engine, s *services.UsersService) {
	grp := r.Group("/users")
	grp.POST("/", routeWrapper[*services.UsersService](s, createOneUser))
	grp.GET("/", routeWrapper[*services.UsersService](s, getManyUsers))
	grp.GET("/profile", routeWrapper[*services.UsersService](s, getUserByEmail))
	grp.GET("/:id", routeWrapper[*services.UsersService](s, getUserById))
	grp.PUT("/:id", routeWrapper[*services.UsersService](s, updateOneUser))
	grp.DELETE("/:id", routeWrapper[*services.UsersService](s, deleteOneUser))
}
