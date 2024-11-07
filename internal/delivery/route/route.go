package route

import (
	"github.com/gin-gonic/gin"
	"personal-projects/templateProjects/internal/delivery"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController delivery.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.GET("/login", c.UserController.Login)
}
