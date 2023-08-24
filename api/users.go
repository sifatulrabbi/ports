package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sifatulrabbi/ports/services"
)

func getUserById(c *gin.Context, s *services.UsersService) {
	res := httpResponse{}
	user, err := s.GetOne(services.UserFilter{Email: "sifatuli.r@gmail.com"})
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
	res := httpResponse{}
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

}

func deleteOneUser(c *gin.Context, s *services.UsersService) {

}

func getManyUsers(c *gin.Context, s *services.UsersService) {

}

func RegisterUsersHandlers(r *gin.Engine, s *services.UsersService) {
	grp := r.Group("/users")
	grp.POST("/", routeWrapper[*services.UsersService](s, createOneUser))
	grp.GET("/", routeWrapper[*services.UsersService](s, getManyUsers))
	grp.GET("/:id", routeWrapper[*services.UsersService](s, getUserById))
	grp.PUT("/:id", routeWrapper[*services.UsersService](s, updateOneUser))
	grp.DELETE("/:id", routeWrapper[*services.UsersService](s, deleteOneUser))
}
