package config

import (
	"github.com/gin-gonic/gin"
	"personal-projects/templateProjects/internal/delivery/http"
	"personal-projects/templateProjects/internal/delivery/http/route"
)

type BootstrapConfig struct {
	App *gin.Engine
}

func Bootstrap(config *BootstrapConfig) {
	userController := http.NewUserController()
	routeConfig := route.RoutesConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
