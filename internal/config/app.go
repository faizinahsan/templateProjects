package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"personal-projects/templateProjects/internal/delivery/http"
	"personal-projects/templateProjects/internal/delivery/http/route"
)

type BootstrapConfig struct {
	App    *gin.Engine
	Config *viper.Viper
	DB     *gorm.DB
	Log    *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	userController := http.NewUserController()
	routeConfig := route.RoutesConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
