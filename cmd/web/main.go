package main

import (
	"fmt"
	"log"
	"personal-projects/templateProjects/internal/config"
	"personal-projects/templateProjects/internal/delivery"
	"personal-projects/templateProjects/internal/delivery/route"
)

func main() {
	app := config.NewGin()
	routes := route.RouteConfig{
		App:            app,
		UserController: delivery.UserController{},
	}
	routes.Setup()
	err := app.Run(fmt.Sprintf("localhost:8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
