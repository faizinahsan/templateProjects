package route

import (
	"github.com/gin-gonic/gin"
	"personal-projects/templateProjects/internal/delivery/http"
)

type RoutesConfig struct {
	App            *gin.Engine
	UserController *http.UserController
}

func (c *RoutesConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RoutesConfig) SetupGuestRoute() {
	c.App.GET("/login", c.UserController.Login)
}
