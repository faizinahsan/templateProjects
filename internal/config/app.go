package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"personal-projects/templateProjects/internal/delivery/http"
	"personal-projects/templateProjects/internal/delivery/http/route"
)

type BootstrapConfig struct {
	App    *gin.Engine
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userController := http.NewUserController()
	routeConfig := route.RoutesConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
