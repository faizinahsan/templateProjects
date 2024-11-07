package main

import (
	"fmt"
	"log"
	"personal-projects/templateProjects/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewGin()
	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
	})
	host := viperConfig.GetString("web.host")
	port := viperConfig.GetString("web.port")
	err := app.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
