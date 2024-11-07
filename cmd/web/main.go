package main

import (
	"fmt"
	"log"
	"personal-projects/templateProjects/internal/config"
)

func main() {
	app := config.NewGin()
	config.Bootstrap(&config.BootstrapConfig{App: app})
	err := app.Run(fmt.Sprintf("localhost:8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
