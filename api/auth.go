package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sifatulrabbi/ports/services"
)

func login(c *gin.Context, service *services.AuthService) {
	res := httpResponse{}
	res.Ok(c)
}

func register(c *gin.Context, service *services.AuthService) {
	res := httpResponse{}
	res.Ok(c)
}

func getAccessToken(c *gin.Context, service *services.AuthService) {
	res := httpResponse{}
	res.Ok(c)
}

func logout(c *gin.Context, service *services.AuthService) {
	res := httpResponse{}
	res.Ok(c)
}

func RegisterAuthHandlers(r *gin.Engine, s *services.AuthService) {
	grp := r.Group("/auth")
	grp.POST("/login", routeWrapper[*services.AuthService](s, login))
	grp.POST("/logout", routeWrapper[*services.AuthService](s, logout))
	grp.POST("/register", routeWrapper[*services.AuthService](s, register))
	grp.GET("/get-access-token", routeWrapper[*services.AuthService](s, getAccessToken))
}
